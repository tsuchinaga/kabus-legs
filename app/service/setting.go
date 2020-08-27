package service

import "gitlab.com/tsuchinaga/kabus-legs/app/repository"

// setting - 設定サービス
type setting struct {
	settingStore repository.SettingStore
	kabuAPI      repository.KabuAPI
}

// SavePassword - パスワードを保存する
func (s *setting) SavePassword(password string) {
	s.settingStore.SetPassword(password)
}

// IsPasswordSet - パスワードが設定済みか
func (s *setting) IsPasswordSet() bool {
	return s.settingStore.GetPassword() != ""
}

// SetIsProd - 本番を向いているかをセットする
func (s *setting) SetIsProd(isProd bool) {
	s.settingStore.SetIsProd(isProd)
}

// IsProd - 本番向きかを返す
func (s *setting) IsProd() bool {
	return s.settingStore.IsProd()
}

// SaveToken - トークンを保存する
func (s *setting) SaveToken(token string) {
	s.settingStore.SetToken(token)
}

// GetToken - ストアからトークンを取得する処理を追加
func (s *setting) GetToken() string {
	return s.settingStore.GetToken()
}

// GetNewToken - 新しいトークンを発行する
func (s *setting) GetNewToken() (string, error) {
	return s.kabuAPI.GetToken()
}
