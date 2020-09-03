package service

import (
	"time"

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
	getToken1        string
	getToken2        error
	registerSymbol   error
	unregisterSymbol error
}

func (t *testKabusAPI) RegisterSymbol(string, value.Exchange) error {
	return t.registerSymbol
}
func (t *testKabusAPI) UnregisterSymbol(string, value.Exchange) error {
	return t.unregisterSymbol
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

type testPriceWebSocket struct {
	start error
	stop  error
}

func (t *testPriceWebSocket) Start() error    { return t.start }
func (t *testPriceWebSocket) Stop() error     { return t.stop }
func (t *testPriceWebSocket) IsStarted() bool { panic("implement me") }

type testTickStore struct {
	addHis []struct {
		price value.Price
		label string
	}
	get []value.Price
}

func (t *testTickStore) Get(value.Symbol, string) []value.Price { return t.get }
func (t *testTickStore) Add(price value.Price, label string) {
	if t.addHis == nil {
		t.addHis = []struct {
			price value.Price
			label string
		}{}
	}
	t.addHis = append(t.addHis, struct {
		price value.Price
		label string
	}{price: price, label: label})
}

type testLegStore struct {
	getN   []value.FourPrice
	get1   []value.FourPrice
	addHis []value.FourPrice
}

func (t *testLegStore) Add(arg value.FourPrice) {
	if t.addHis == nil {
		t.addHis = []value.FourPrice{}
	}
	t.addHis = append(t.addHis, arg)
}
func (t *testLegStore) Get(_ value.Symbol, legPeriod int) []value.FourPrice {
	if legPeriod == 1 {
		return t.get1
	}
	return t.getN
}

type testClock struct{ now time.Time }

func (t *testClock) Now() time.Time { return t.now }
