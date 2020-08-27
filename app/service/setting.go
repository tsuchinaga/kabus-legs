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
