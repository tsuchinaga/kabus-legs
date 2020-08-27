package kabu

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

type testSettingStore struct {
	getPassword string
}

func (t *testSettingStore) IsPasswordSet() bool { panic("implement me") }
func (t *testSettingStore) GetPassword() string { return t.getPassword }
func (t *testSettingStore) SetPassword(string)  { panic("implement me") }
func (t *testSettingStore) GetToken() string    { panic("implement me") }
func (t *testSettingStore) SetToken(string)     { panic("implement me") }
func (t *testSettingStore) IsProd() bool        { panic("implement me") }
func (t *testSettingStore) SetIsProd(bool)      { panic("implement me") }

type testTokenRequester struct {
	ret1 *kabus.TokenResponse
	ret2 error
}

func (t *testTokenRequester) Exec(kabus.TokenRequest) (*kabus.TokenResponse, error) {
	return t.ret1, t.ret2
}