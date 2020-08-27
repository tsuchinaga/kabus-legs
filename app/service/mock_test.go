package service

type testSettingStore struct {
	setPasswordHis []string
}

func (t *testSettingStore) IsPasswordSet() bool { panic("implement me") }
func (t *testSettingStore) GetPassword() string { panic("implement me") }
func (t *testSettingStore) SetPassword(password string) {
	if t.setPasswordHis == nil {
		t.setPasswordHis = []string{}
	}
	t.setPasswordHis = append(t.setPasswordHis, password)
}
func (t *testSettingStore) GetToken() string { panic("implement me") }
func (t *testSettingStore) SetToken(string)  { panic("implement me") }
func (t *testSettingStore) IsProd() bool     { panic("implement me") }
func (t *testSettingStore) SetIsProd(bool)   { panic("implement me") }
