package usecase

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"
)

func Test_setting_GetSavedToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		getToken string
		want1    string
		want2    error
	}{
		{name: "保存しているトークンを取り出す", getToken: "token", want1: "token", want2: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingService := &testSettingService{getToken: test.getToken}
			usecase := &setting{settingService: settingService}
			got1, got2 := usecase.GetToken()
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_setting_GetNewToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		getNewToken1 string
		getNewToken2 error
		want1        string
		want2        error
	}{
		{name: "新しいトークンを取得する", getNewToken1: "token", getNewToken2: nil, want1: "token", want2: nil},
		{name: "新しいトークン取得でエラーがあったらエラーを返す",
			getNewToken1: "", getNewToken2: app.APIRequestError, want1: "", want2: app.APIRequestError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingService := &testSettingService{getNewToken1: test.getNewToken1, getNewToken2: test.getNewToken2}
			usecase := &setting{settingService: settingService}
			got1, got2 := usecase.GetNewToken()
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
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
		{name: "トークンをセットする", arg: "token", want: []string{"token"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingService := &testSettingService{}
			usecase := &setting{settingService: settingService}
			usecase.SaveToken(test.arg)
			if !reflect.DeepEqual(test.want, settingService.saveTokenHis) {
				t.Errorf("%s error\nwant: %+v\nhistory: %+v\n", t.Name(), test.want, settingService.saveTokenHis)
			}
		})
	}
}

func Test_setting_SavePassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{name: "パスワードを保存する", arg: "password", want: []string{"password"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			settingService := &testSettingService{}
			usecase := &setting{settingService: settingService}
			usecase.SavePassword(test.arg)
			if !reflect.DeepEqual(test.want, settingService.savePasswordHis) {
				t.Errorf("%s error\nwant: %+v\nhistory: %+v\n", t.Name(), test.want, settingService.savePasswordHis)
			}
		})
	}
}
