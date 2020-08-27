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
				time.Sleep(100 * time.Millisecond)
				store.password = test.password
			}()
			got := store.IsPasswordSet()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
