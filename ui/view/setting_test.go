package view

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_SettingStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg1 value.SettingStatus
		arg2 error
		want string
	}{
		{name: "エラーがあればエラーを表示",
			arg1: value.SettingStatus{}, arg2: errors.New("error message"),
			want: "エラーが発生しました(error message)"},
		{name: "設定済みの本番",
			arg1: value.SettingStatus{IsPasswordSet: true, IsProd: true}, arg2: nil,
			want: "パスワード: 設定済み, 環境: 本番"},
		{name: "未設定の検証",
			arg1: value.SettingStatus{IsPasswordSet: false, IsProd: false}, arg2: nil,
			want: "パスワード: 未設定, 環境: 検証"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := SettingStatus(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
