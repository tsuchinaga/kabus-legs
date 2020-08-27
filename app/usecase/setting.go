package usecase

import "gitlab.com/tsuchinaga/kabus-legs/app/service"

type setting struct {
	settingService service.Setting
}

// GetToken - 保存されたトークンを取り出す
func (u *setting) GetToken() (string, error) {
	return u.settingService.GetToken(), nil
}
