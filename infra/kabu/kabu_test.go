package kabu

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func Test_kabu_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		getPassword string
		tokenRet1   *kabus.TokenResponse
		tokenRet2   error
		want1       string
		want2       error
	}{
		{name: "errorが返されたらAPIRequestErrorでラップして返す",
			getPassword: "password", tokenRet1: nil, tokenRet2: errors.New("foo bar"),
			want1: "", want2: app.APIRequestError},
		{name: "エラーでなければ取得したトークンを返す",
			getPassword: "password", tokenRet1: &kabus.TokenResponse{Token: "token"}, tokenRet2: nil,
			want1: "token", want2: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{getPassword: test.getPassword}
			tokenRequester := &testTokenRequester{ret1: test.tokenRet1, ret2: test.tokenRet2}
			k := &kabu{settingStore: settingStore, tokenRequester: tokenRequester}
			got1, got2 := k.GetToken()
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
