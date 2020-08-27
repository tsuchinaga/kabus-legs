package service

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
