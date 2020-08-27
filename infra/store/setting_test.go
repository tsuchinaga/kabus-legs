package store

import (
	"reflect"
	"testing"
	"time"
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
