package kabu

import (
	"fmt"
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

var (
	priceWS    repository.PriceWebSocket
	priceWSMtx sync.Mutex
)

// GetPrice - 新しい現値WebSocketを生成する
func GetPrice(settingStore repository.SettingStore, f func(value.Price) error) repository.PriceWebSocket {
	priceWSMtx.Lock()
	defer priceWSMtx.Unlock()

	// 既に存在しているなら存在していたのを返す
	if priceWS != nil {
		return priceWS
	}

	priceWS = &price{
		settingStore: settingStore,
		priceWSRequester: kabus.NewWSRequester(settingStore.IsProd(), func(message kabus.PriceMessage) error {
			if err := f(value.Price{
				SymbolCode: message.Symbol,
				Exchange:   convertExchange(message.Exchange),
				Price:      message.CurrentPrice,
				Time:       message.CurrentPriceTime,
			}); err != nil {
				return err
			}
			return nil
		}),
	}

	return priceWS
}

// price - 現値WebSocket
type price struct {
	settingStore     repository.SettingStore
	priceWSRequester PriceWSRequester
	started          bool
	mtx              sync.Mutex
}

// Start - WebSocketの接続を開始する
func (p *price) Start() error {
	p.mtx.Lock()
	if p.started == true {
		p.mtx.Unlock()
		return app.WebSocketIsStartedError
	}

	p.started = true
	p.mtx.Unlock()

	if err := p.priceWSRequester.Open(); err != nil {
		return fmt.Errorf("%v: %w", err, app.APIRequestError)
	}
	return nil
}

// Stop - WebSocketの接続を停止する
func (p *price) Stop() error {
	p.mtx.Lock()
	if p.started == false {
		p.mtx.Unlock()
		return app.WebSocketIsStoppedError
	}

	p.started = false
	p.mtx.Unlock()

	defer func() {
		priceWSMtx.Lock()
		priceWS = nil
		priceWSMtx.Unlock()
	}()

	if err := p.priceWSRequester.Close(); err != nil {
		return fmt.Errorf("%v: %w", err, app.APIRequestError)
	}
	return nil
}

// IsStarted - WebSocketが接続されているか
func (p *price) IsStarted() bool {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	return p.started
}
