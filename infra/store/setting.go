package store

import "sync"

// setting - 設定ストア
type setting struct {
	password string // パスワード
	token    string // トークン
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
