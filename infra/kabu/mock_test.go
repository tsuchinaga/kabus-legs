package kabu

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

type testSettingStore struct {
	getPassword string
	isProd      bool
	token       string
}

func (t *testSettingStore) IsPasswordSet() bool { panic("implement me") }
func (t *testSettingStore) GetPassword() string { return t.getPassword }
func (t *testSettingStore) SetPassword(string)  { panic("implement me") }
func (t *testSettingStore) GetToken() string    { return t.token }
func (t *testSettingStore) SetToken(string)     { panic("implement me") }
func (t *testSettingStore) IsProd() bool        { return t.isProd }
func (t *testSettingStore) SetIsProd(bool)      { panic("implement me") }

type testTokenRequester struct {
	ret1 *kabus.TokenResponse
	ret2 error
}

func (t *testTokenRequester) Exec(kabus.TokenRequest) (*kabus.TokenResponse, error) {
	return t.ret1, t.ret2
}

type testRegisterRequester struct {
	exec1 *kabus.RegisterResponse
	exec2 error
}

func (t *testRegisterRequester) Exec(kabus.RegisterRequest) (*kabus.RegisterResponse, error) {
	return t.exec1, t.exec2
}

type testUnregisterRequester struct {
	exec1 *kabus.UnregisterResponse
	exec2 error
}

func (t *testUnregisterRequester) Exec(kabus.UnregisterRequest) (*kabus.UnregisterResponse, error) {
	return t.exec1, t.exec2
}
