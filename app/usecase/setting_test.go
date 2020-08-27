package usecase

import (
	"errors"
	"reflect"
	"testing"
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
