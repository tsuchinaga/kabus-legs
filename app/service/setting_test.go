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
