package service

import (
	"reflect"
	"testing"
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
