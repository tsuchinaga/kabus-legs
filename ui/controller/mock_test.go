package controller

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

type testSettingUseCase struct {
	getToken1    string
	getToken2    error
	getNewToken1 string
	getNewToken2 error
}

func (t *testSettingUseCase) GetToken() (string, error)    { return t.getToken1, t.getToken2 }
func (t *testSettingUseCase) GetNewToken() (string, error) { return t.getNewToken1, t.getNewToken2 }
func (t *testSettingUseCase) SaveToken(string)             {}
func (t *testSettingUseCase) SavePassword(string)          {}
func (t *testSettingUseCase) SetIsProd(bool)               {}
func (t *testSettingUseCase) GetSettingStatus() (value.SettingStatus, error) {
	return value.SettingStatus{}, nil
}

type testSymbolLegUseCase struct {
	getAll1 []value.SymbolLeg
	getAll2 error
}

func (t *testSymbolLegUseCase) GetAll() ([]value.SymbolLeg, error) { return t.getAll1, t.getAll2 }
func (t *testSymbolLegUseCase) Register(string, string, int) error { panic("implement me") }
func (t *testSymbolLegUseCase) Unregister(int) error               { panic("implement me") }

type testSettingView struct {
	settingStatus string
}

func (t *testSettingView) SettingStatus(value.SettingStatus, error) string {
	return t.settingStatus
}
