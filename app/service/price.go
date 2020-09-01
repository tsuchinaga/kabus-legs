package service

import "gitlab.com/tsuchinaga/kabus-legs/app/repository"

// price - 価格サービスのインターフェース
type Price interface {
	StartWebSocket() error
	StopWebSocket() error
}

// NewPriceWebSocket - 新しい価格サービスを生成する
func NewPriceWebSocket(priceWebSocket repository.PriceWebSocket) Price {
	return &price{priceWebSocket: priceWebSocket}
}

// price - 価格サービス
type price struct {
	priceWebSocket repository.PriceWebSocket
}

// StartWebSocket - WebSocketを開始する
func (p *price) StartWebSocket() error {
	return p.priceWebSocket.Start()
}

// StopWebSocket - WebSocketを停止する
func (p *price) StopWebSocket() error {
	return p.priceWebSocket.Stop()
}
