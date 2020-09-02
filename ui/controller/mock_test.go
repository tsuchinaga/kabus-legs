package controller

import (
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

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
	getAll1    []value.SymbolLeg
	getAll2    error
	register   error
	unregister error
}

func (t *testSymbolLegUseCase) GetAll() ([]value.SymbolLeg, error) { return t.getAll1, t.getAll2 }
func (t *testSymbolLegUseCase) Register(string, string, int) error { return t.register }
func (t *testSymbolLegUseCase) Unregister(int) error               { return t.unregister }

type testSettingView struct {
	settingStatus string
}

func (t *testSettingView) SettingStatus(value.SettingStatus, error) string {
	return t.settingStatus
}

type testPriceUseCase struct {
	startGetPrice   error
	startReturnWait time.Duration
	stopGetPrice    error
}

func (t *testPriceUseCase) StartGetPrice() error {
	time.Sleep(t.startReturnWait)
	return t.startGetPrice
}
func (t *testPriceUseCase) StopGetPrice() error { return t.stopGetPrice }
