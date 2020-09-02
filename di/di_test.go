package di

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/ui/controller"
	"gitlab.com/tsuchinaga/kabus-legs/ui/view"

	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"
	"gitlab.com/tsuchinaga/kabus-legs/infra/kabu"
	"gitlab.com/tsuchinaga/kabus-legs/infra/store"
)

func Test_NewSettingController(t *testing.T) {
	t.Parallel()
	want := controller.NewSetting(
		NewSettingUseCase(),
		view.NewSetting(),
	)

	got := NewSettingController()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

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

func Test_NewTokenController(t *testing.T) {
	t.Parallel()
	want := controller.NewToken(NewSettingUseCase())
	got := NewTokenController()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_NewSymbolController(t *testing.T) {
	t.Parallel()
	want := controller.NewSymbol(usecase.NewSymbolLeg(service.NewSymbol(store.GetSymbol(), kabu.NewKabuAPI(store.GetSetting()))))
	got := NewSymbolController()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_NewPriceController(t *testing.T) {
	t.Parallel()
	want := controller.NewPrice(usecase.NewPrice(service.NewPriceWebSocket(kabu.GetPrice(store.GetSetting(), nil))))
	got := NewPriceController()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
