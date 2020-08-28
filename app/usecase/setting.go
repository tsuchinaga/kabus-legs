package usecase

import "gitlab.com/tsuchinaga/kabus-legs/app/service"

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
