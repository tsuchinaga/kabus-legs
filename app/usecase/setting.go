package usecase

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

type setting struct {
	settingService service.Setting
}

// GetToken - 保存されたトークンを取り出す
func (u *setting) GetToken() (string, error) {
	return u.settingService.GetToken(), nil
}

// GetNewToken - APIを使って新しいトークンを取得する
func (u *setting) GetNewToken() (string, error) {
	return u.settingService.GetNewToken()
}

// SaveToken - 引数のトークンを保存する
func (u *setting) SaveToken(token string) {
	u.settingService.SaveToken(token)
}

// SavePassword - 引数のパスワードを保存する
func (u *setting) SavePassword(password string) {
	u.settingService.SavePassword(password)
}

// SetIsProd - 本番向きか検証向きかを設定する
func (u *setting) SetIsProd(isProd bool) {
	u.settingService.SetIsProd(isProd)
}

// GetSettingStatus - 設定の状況を取得する
func (u *setting) GetSettingStatus() (value.SettingStatus, error) {
	return value.SettingStatus{IsPasswordSet: u.settingService.IsPasswordSet(), IsProd: u.settingService.IsProd()}, nil
}
