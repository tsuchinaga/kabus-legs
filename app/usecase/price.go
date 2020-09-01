package usecase

import "gitlab.com/tsuchinaga/kabus-legs/app/service"

// Price - 価格ユースケースのインターフェース
type Price interface {
	StartGetPrice() error
	StopGetPrice() error
}

// NewPrice - 価格ユースケースの生成
func NewPrice(priceService service.Price) Price {
	return &price{priceService: priceService}
}

// price - 価格ユースケース
type price struct {
	priceService service.Price
}

// StartGetPrice - 価格取得の開始
func (p *price) StartGetPrice() error {
	return p.priceService.StartWebSocket()
}

// StopGetPrice - 価格取得の停止
func (p *price) StopGetPrice() error {
	return p.priceService.StopWebSocket()
}
