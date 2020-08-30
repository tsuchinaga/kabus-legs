package kabu

import (
	"fmt"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-legs/app"
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

// NewKabuAPI - kabus apiを呼び出す処理群を生成する
func NewKabuAPI(settingStore repository.SettingStore) repository.KabuAPI {
	return &kabu{
		settingStore:        settingStore,
		tokenRequester:      kabus.NewTokenRequester(settingStore.IsProd()),
		registerRequester:   kabus.NewRegisterRequester(settingStore.GetToken(), settingStore.IsProd()),
		unregisterRequester: kabus.NewUnregisterRequester(settingStore.GetToken(), settingStore.IsProd()),
	}
}

// kabu - kabus apiを呼び出す処理をまとめた構造体
type kabu struct {
	settingStore        repository.SettingStore
	tokenRequester      TokenRequester
	registerRequester   RegisterRequester
	unregisterRequester UnregisterRequester
}

// GetToken - トークンの取得
func (k *kabu) GetToken() (string, error) {
	res, err := k.tokenRequester.Exec(kabus.TokenRequest{APIPassword: k.settingStore.GetPassword()})
	if err != nil {
		return "", fmt.Errorf("%v: %w", err, app.APIRequestError)
	}
	return res.Token, nil
}

// RegisterSymbol - 銘柄登録に引数の銘柄を渡す
func (k *kabu) RegisterSymbol(symbolCode string, exchange value.Exchange) error {
	_, err := k.registerRequester.Exec(kabus.RegisterRequest{
		Symbols: []kabus.RegistSymbol{
			{Symbol: symbolCode, Exchange: toKabusExchange(exchange)},
		},
	})
	if err != nil {
		return fmt.Errorf("%v: %w", err, app.APIRequestError)
	}
	return nil
}

// toKabusExchange - kabusのExchangeに変換する
func toKabusExchange(exchange value.Exchange) kabus.Exchange {
	switch exchange {
	case value.ExchangeT:
		return kabus.ExchangeToushou
	case value.ExchangeM:
		return kabus.ExchangeMeishou
	case value.ExchangeF:
		return kabus.ExchangeFukushou
	case value.ExchangeS:
		return kabus.ExchangeSatsushou
	}
	return kabus.ExchangeUnspecified
}
