package store

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

var (
	settingStore    repository.SettingStore
	settingStoreMtx sync.Mutex
)

// GetSetting - 設定のストアを取り出す
func GetSetting() repository.SettingStore {
	settingStoreMtx.Lock()
	defer settingStoreMtx.Unlock()

	if settingStore == nil {
		settingStore = &setting{}
	}
	return settingStore
}

// setting - 設定ストア
type setting struct {
	password string // パスワード
	token    string // トークン
	isProd   bool   // 本番か
	mtx      sync.Mutex
}

// IsPasswordSet - パスワードが設定済みか
func (s *setting) IsPasswordSet() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.password != ""
}

// SetPassword - パスワードをセットする
func (s *setting) SetPassword(password string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.password = password
}

// GetToken - トークンを取得する
func (s *setting) GetToken() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.token
}

// SetToken - トークンをセットする
func (s *setting) SetToken(token string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.token = token
}

// IsProd - 本番向きかを返す
func (s *setting) IsProd() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.isProd
}

// SetIsProd - 本番かをセットする
func (s *setting) SetIsProd(isProd bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.isProd = isProd
}

// GetPassword - パスワードを取得する
func (s *setting) GetPassword() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.password
}
