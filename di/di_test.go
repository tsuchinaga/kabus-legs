package di

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"
	"gitlab.com/tsuchinaga/kabus-legs/infra/kabu"
	"gitlab.com/tsuchinaga/kabus-legs/infra/store"
)

func Test_NewSettingUseCase(t *testing.T) {
	t.Parallel()
	want := usecase.NewSetting(
		service.NewSetting(
			store.GetSetting(),
			kabu.NewKabuAPI(
				store.GetSetting(),
			),
		),
	)

	got := NewSettingUseCase()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
