package service

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"
)

func Test_setting_SavePassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{name: "passwordをstoreにsetできる", arg: "password", want: []string{"password"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{}
			service := &setting{settingStore: settingStore}
			service.SavePassword(test.arg)
			if !reflect.DeepEqual(test.want, settingStore.setPasswordHis) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, settingStore.setPasswordHis)
			}
		})
	}
}

func Test_setting_IsSetPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		getPassword string
		want        bool
	}{
		{name: "パスワードが空文字なら未設定", getPassword: "", want: false},
		{name: "パスワードが空文字じゃなければ設定済み", getPassword: "password", want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{getPassword: test.getPassword}
			service := &setting{settingStore: settingStore}
			got := service.IsPasswordSet()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SetIsProd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  bool
		want []bool
	}{
		{name: "本番であることを設定できる", arg: true, want: []bool{true}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{}
			service := &setting{settingStore: settingStore}
			service.SetIsProd(test.arg)
			if !reflect.DeepEqual(test.want, settingStore.setIsProdHis) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, settingStore.setIsProdHis)
			}
		})
	}
}

func Test_setting_IsProd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		isProd bool
		want   bool
	}{
		{name: "本番を向いているかを返せる", isProd: true, want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{isProd: test.isProd}
			service := &setting{settingStore: settingStore}
			got := service.IsProd()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SaveToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{name: "tokenをセットできる", arg: "token", want: []string{"token"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{}
			service := &setting{settingStore: settingStore}
			service.SaveToken(test.arg)
			if !reflect.DeepEqual(test.want, settingStore.setTokenHis) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, settingStore.setTokenHis)
			}
		})
	}
}

func Test_setting_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		getToken string
		want     string
	}{
		{name: "tokenを取得できる", getToken: "token", want: "token"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingStore := &testSettingStore{getToken: test.getToken}
			service := &setting{settingStore: settingStore}
			got := service.GetToken()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_Test_setting_GetNewToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		want1     string
		want2     error
	}{
		{name: "新しいトークンを取得する",
			getToken1: "token", getToken2: nil,
			want1: "token", want2: nil},
		{name: "トークン取得に失敗",
			getToken1: "", getToken2: app.APIRequestError,
			want1: "", want2: app.APIRequestError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &setting{kabuAPI: &testKabusAPI{getToken1: test.getToken1, getToken2: test.getToken2}}
			got1, got2 := service.GetNewToken()
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
