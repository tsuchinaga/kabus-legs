package usecase

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// Setting - 設定ユースケースのインターフェース
type Setting interface {
	GetToken() (string, error)
	GetNewToken() (string, error)
	SaveToken(token string)
	SavePassword(password string)
	SetIsProd(isProd bool)
	GetSettingStatus() (value.SettingStatus, error)
}

// setting - 設定ユースケースの生成
func NewSetting(settingService service.Setting) Setting {
	return &setting{settingService: settingService}
}

// setting - 設定ユースケース
type setting struct {
	settingService service.Setting
}

// GetToken - 保存されたトークンを取り出す
func (u *setting) GetToken() (string, error) {
	return u.settingService.GetToken(), nil
}

// GetNewToken - APIを使って新しいトークンを取得する
func (u *setting) GetNewToken() (string, error) {
	token, err := u.settingService.GetNewToken()
	if err != nil {
		return "", err
	}
	u.settingService.SaveToken(token)
	return token, nil
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
