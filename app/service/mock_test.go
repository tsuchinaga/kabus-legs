package service

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

type testSettingStore struct {
	getPassword    string
	setPasswordHis []string
	isProd         bool
	setIsProdHis   []bool
	getToken       string
	setTokenHis    []string
}

func (t *testSettingStore) GetPassword() string { return t.getPassword }
func (t *testSettingStore) SetPassword(password string) {
	if t.setPasswordHis == nil {
		t.setPasswordHis = []string{}
	}
	t.setPasswordHis = append(t.setPasswordHis, password)
}
func (t *testSettingStore) GetToken() string { return t.getToken }
func (t *testSettingStore) SetToken(token string) {
	if t.setTokenHis == nil {
		t.setTokenHis = []string{}
	}
	t.setTokenHis = append(t.setTokenHis, token)
}
func (t *testSettingStore) IsProd() bool { return t.isProd }
func (t *testSettingStore) SetIsProd(isProd bool) {
	if t.setIsProdHis == nil {
		t.setIsProdHis = []bool{}
	}
	t.setIsProdHis = append(t.setIsProdHis, isProd)
}

type testKabusAPI struct {
	getToken1      string
	getToken2      error
	registerSymbol error
}

func (t *testKabusAPI) RegisterSymbol(string, value.Exchange) error {
	return t.registerSymbol
}
func (t *testKabusAPI) UnregisterSymbol(string, value.Exchange) error {
	panic("implement me")
}
func (t *testKabusAPI) GetToken() (string, error) { return t.getToken1, t.getToken2 }

type testSymbolStore struct {
	getAll           []value.SymbolLeg
	addHis           []value.SymbolLeg
	deleteByIndexHis []int
}

func (t *testSymbolStore) IsExists(value.SymbolLeg) bool { panic("implement me") }
func (t *testSymbolStore) Add(symbolLeg value.SymbolLeg) {
	if t.addHis == nil {
		t.addHis = []value.SymbolLeg{}
	}
	t.addHis = append(t.addHis, symbolLeg)
}
func (t *testSymbolStore) GetAll() []value.SymbolLeg { return t.getAll }
func (t *testSymbolStore) DeleteByIndex(i int) {
	if t.deleteByIndexHis == nil {
		t.deleteByIndexHis = []int{}
	}
	t.deleteByIndexHis = append(t.deleteByIndexHis, i)
}
