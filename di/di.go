package di

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
	"gitlab.com/tsuchinaga/kabus-legs/infra/kabu"
	"gitlab.com/tsuchinaga/kabus-legs/infra/store"
	"gitlab.com/tsuchinaga/kabus-legs/ui/controller"
	"gitlab.com/tsuchinaga/kabus-legs/ui/view"
)

// NewSettingController - DI済みの設定コントローラを生成する
func NewSettingController() controller.Setting {
	return controller.NewSetting(
		NewSettingUseCase(),
		view.NewSetting(),
	)
}

// NewTokenController - DI済みのトークンコントローラを生成する
func NewTokenController() controller.Token {
	return controller.NewToken(NewSettingUseCase())
}

// NewSymbolController - DI済みの銘柄コントローラを生成する
func NewSymbolController() controller.Symbol {
	return controller.NewSymbol(
		usecase.NewSymbolLeg(
			service.NewSymbol(
				store.GetSymbol(),
				kabu.NewKabuAPI(
					store.GetSetting(),
				),
			),
		),
	)
}

// NewPriceController - DI済みの価格コントローラの生成
func NewPriceController() controller.Price {
	f := func(price value.Price) error {
		tickService := service.NewTick(store.GetTick())
		tickService.SavePrice(price)
		return nil
	}

	return controller.NewPrice(
		usecase.NewPrice(
			service.NewPriceWebSocket(
				kabu.GetPrice(store.GetSetting(), f))))
}

// NewSettingUseCase - DI済みの設定ユースケースを生成する
func NewSettingUseCase() usecase.Setting {
	return usecase.NewSetting(
		service.NewSetting(
			store.GetSetting(),
			kabu.NewKabuAPI(
				store.GetSetting(),
			),
		),
	)
}
