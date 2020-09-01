package usecase

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

type testSettingService struct {
	savePasswordHis []string
	isPasswordSet   bool
	setIsProdHis    []bool
	isProd          bool
	saveTokenHis    []string
	getToken        string
	getNewToken1    string
	getNewToken2    error
}

func (t *testSettingService) SavePassword(password string) {
	if t.savePasswordHis == nil {
		t.savePasswordHis = []string{}
	}
	t.savePasswordHis = append(t.savePasswordHis, password)
}
func (t *testSettingService) IsPasswordSet() bool { return t.isPasswordSet }
func (t *testSettingService) SetIsProd(isProd bool) {
	if t.setIsProdHis == nil {
		t.setIsProdHis = []bool{}
	}
	t.setIsProdHis = append(t.setIsProdHis, isProd)
}
func (t *testSettingService) IsProd() bool { return t.isProd }
func (t *testSettingService) SaveToken(token string) {
	if t.saveTokenHis == nil {
		t.saveTokenHis = []string{}
	}
	t.saveTokenHis = append(t.saveTokenHis, token)
}
func (t *testSettingService) GetToken() string             { return t.getToken }
func (t *testSettingService) GetNewToken() (string, error) { return t.getNewToken1, t.getNewToken2 }

type testSymbolService struct {
	getAll         []value.SymbolLeg
	sendRegister   error
	getByIndex1    value.SymbolLeg
	getByIndex2    error
	sendUnregister error
	getBySymbol    []value.SymbolLeg
}

func (t *testSymbolService) GetBySymbol(string, value.Exchange) []value.SymbolLeg {
	return t.getBySymbol
}
func (t *testSymbolService) GetByIndex(int) (value.SymbolLeg, error) {
	return t.getByIndex1, t.getByIndex2
}
func (t *testSymbolService) GetAll() []value.SymbolLeg                   { return t.getAll }
func (t *testSymbolService) AddSymbol(value.SymbolLeg)                   {}
func (t *testSymbolService) DeleteSymbolByIndex(int)                     {}
func (t *testSymbolService) SendRegister(string, value.Exchange) error   { return t.sendRegister }
func (t *testSymbolService) SendUnregister(string, value.Exchange) error { return t.sendUnregister }

type testPriceService struct {
	startWebSocket error
	stopWebSocket  error
}

func (t *testPriceService) StartWebSocket() error { return t.startWebSocket }
func (t *testPriceService) StopWebSocket() error  { return t.stopWebSocket }
