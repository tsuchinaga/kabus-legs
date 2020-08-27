package store

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

func Test_setting_IsPasswordSet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{name: "passwordが空文字ならfalse", password: "", want: false},
		{name: "passwordが空文字でなければtrue", password: "password", want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{}
			store.mtx.Lock()
			go func() {
				defer store.mtx.Unlock()
				time.Sleep(1 * time.Second)
				store.password = test.password
			}()
			got := store.IsPasswordSet()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SetPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "パスワードに任意の文字列を設定できる", arg: "password", want: "password"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{}
			store.SetPassword(test.arg)
			if !reflect.DeepEqual(test.want, store.password) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, store.password)
			}
		})
	}
}

func Test_setting_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token string
		want  string
	}{
		{name: "tokenを返す", token: "token", want: "token"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{token: test.token}
			got := store.GetToken()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "tokenにセットできる", arg: "token", want: "token"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{}
			store.SetToken(test.arg)
			if !reflect.DeepEqual(test.want, store.token) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, store.token)
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
		{name: "storeからisProdの値が取れる", isProd: true, want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{isProd: test.isProd}
			got := store.IsProd()
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
		want bool
	}{
		{name: "本番かを設定できる", arg: true, want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &setting{}
			store.SetIsProd(test.arg)
			if !reflect.DeepEqual(test.want, store.isProd) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, store.isProd)
			}
		})
	}
}

func Test_GetSetting(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want repository.SettingStore
	}{
		{name: "設定ストアを取得できる", want: &setting{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := GetSetting()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
