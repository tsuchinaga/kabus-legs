package usecase

type testSettingService struct {
	getToken     string
	getNewToken1 string
	getNewToken2 error
}

func (t *testSettingService) SavePassword(string)          { panic("implement me") }
func (t *testSettingService) IsPasswordSet() bool          { panic("implement me") }
func (t *testSettingService) SetIsProd(bool)               { panic("implement me") }
func (t *testSettingService) IsProd() bool                 { panic("implement me") }
func (t *testSettingService) SaveToken(string)             { panic("implement me") }
func (t *testSettingService) GetToken() string             { return t.getToken }
func (t *testSettingService) GetNewToken() (string, error) { return t.getNewToken1, t.getNewToken2 }
