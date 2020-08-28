package usecase

type testSettingService struct {
	savePasswordHis []string
	setIsProdHis    []bool
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
func (t *testSettingService) IsPasswordSet() bool { panic("implement me") }
func (t *testSettingService) SetIsProd(isProd bool) {
	if t.setIsProdHis == nil {
		t.setIsProdHis = []bool{}
	}
	t.setIsProdHis = append(t.setIsProdHis, isProd)
}
func (t *testSettingService) IsProd() bool { panic("implement me") }
func (t *testSettingService) SaveToken(token string) {
	if t.saveTokenHis == nil {
		t.saveTokenHis = []string{}
	}
	t.saveTokenHis = append(t.saveTokenHis, token)
}
func (t *testSettingService) GetToken() string             { return t.getToken }
func (t *testSettingService) GetNewToken() (string, error) { return t.getNewToken1, t.getNewToken2 }
