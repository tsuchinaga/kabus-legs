package kabu

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

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

func Test_NewKabuAPI(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		settingStore repository.SettingStore
		want         repository.KabuAPI
	}{
		{name: "本番向き", settingStore: &testSettingStore{isProd: true, token: "token"},
			want: &kabu{
				settingStore:        &testSettingStore{isProd: true, token: "token"},
				tokenRequester:      kabus.NewTokenRequester(true),
				registerRequester:   kabus.NewRegisterRequester("token", true),
				unregisterRequester: kabus.NewUnregisterRequester("token", true),
			}},
		{name: "検証向き", settingStore: &testSettingStore{isProd: false, token: "token"},
			want: &kabu{
				settingStore:        &testSettingStore{isProd: false, token: "token"},
				tokenRequester:      kabus.NewTokenRequester(false),
				registerRequester:   kabus.NewRegisterRequester("token", false),
				unregisterRequester: kabus.NewUnregisterRequester("token", false),
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewKabuAPI(test.settingStore)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_toKabusExchange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		exchange value.Exchange
		want     kabus.Exchange
	}{
		{name: "Tを変換できる", exchange: value.ExchangeT, want: kabus.ExchangeToushou},
		{name: "Mを変換できる", exchange: value.ExchangeM, want: kabus.ExchangeMeishou},
		{name: "Fを変換できる", exchange: value.ExchangeF, want: kabus.ExchangeFukushou},
		{name: "Mを変換できる", exchange: value.ExchangeS, want: kabus.ExchangeSatsushou},
		{name: "指定なしを変換できる", exchange: value.ExchangeUnspecified, want: kabus.ExchangeUnspecified},
		{name: "想定外なのは全部指定なしになる", exchange: value.Exchange("foo"), want: kabus.ExchangeUnspecified},
	}

	for _, test := range tests {
		test := test
		got := toKabusExchange(test.exchange)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_kabu_RegisterSymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		exec1 *kabus.RegisterResponse
		exec2 error
		arg1  string
		arg2  value.Exchange
		want  error
	}{
		{name: "errorが変えされたらラップして返す",
			exec2: errors.New("error message"),
			want:  app.APIRequestError},
		{name: "errorがなければnilを返す",
			want: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			registerRequester := &testRegisterRequester{exec1: test.exec1, exec2: test.exec2}
			got := (&kabu{registerRequester: registerRequester}).RegisterSymbol(test.arg1, test.arg2)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_kabu_UnregisterSymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		exec1 *kabus.UnregisterResponse
		exec2 error
		arg1  string
		arg2  value.Exchange
		want  error
	}{
		{name: "errorが変えされたらラップして返す",
			exec2: errors.New("error message"),
			want:  app.APIRequestError},
		{name: "errorがなければnilを返す",
			want: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			unregisterRequester := &testUnregisterRequester{exec1: test.exec1, exec2: test.exec2}
			got := (&kabu{unregisterRequester: unregisterRequester}).UnregisterSymbol(test.arg1, test.arg2)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
