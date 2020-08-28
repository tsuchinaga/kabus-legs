package di

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"
	"gitlab.com/tsuchinaga/kabus-legs/infra/kabu"
	"gitlab.com/tsuchinaga/kabus-legs/infra/store"
)

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
