package controller

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

type testSettingUseCase struct{}

func (t *testSettingUseCase) GetToken() (string, error)    { panic("implement me") }
func (t *testSettingUseCase) GetNewToken() (string, error) { panic("implement me") }
func (t *testSettingUseCase) SaveToken(string)             { panic("implement me") }
func (t *testSettingUseCase) SavePassword(string)          {}
func (t *testSettingUseCase) SetIsProd(bool)               {}
func (t *testSettingUseCase) GetSettingStatus() (value.SettingStatus, error) {
	return value.SettingStatus{}, nil
}

type testSettingView struct {
	settingStatus string
}

func (t *testSettingView) SettingStatus(value.SettingStatus, error) string {
	return t.settingStatus
}
