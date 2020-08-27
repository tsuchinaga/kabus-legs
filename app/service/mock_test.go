package service

type testSettingStore struct {
	getPassword    string
	setPasswordHis []string
	setIsProdHis   []bool
}

func (t *testSettingStore) IsPasswordSet() bool { panic("implement me") }
func (t *testSettingStore) GetPassword() string { return t.getPassword }
func (t *testSettingStore) SetPassword(password string) {
	if t.setPasswordHis == nil {
		t.setPasswordHis = []string{}
	}
	t.setPasswordHis = append(t.setPasswordHis, password)
}
func (t *testSettingStore) GetToken() string { panic("implement me") }
func (t *testSettingStore) SetToken(string)  { panic("implement me") }
func (t *testSettingStore) IsProd() bool     { panic("implement me") }
func (t *testSettingStore) SetIsProd(isProd bool) {
	if t.setIsProdHis == nil {
		t.setIsProdHis = []bool{}
	}
	t.setIsProdHis = append(t.setIsProdHis, isProd)
}
