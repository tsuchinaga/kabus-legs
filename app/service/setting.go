package service

import "gitlab.com/tsuchinaga/kabus-legs/app/repository"

// setting - 設定サービス
type setting struct {
	settingStore repository.SettingStore
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
