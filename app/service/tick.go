package service

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// tick - ティックサービスのインターフェース
type Tick interface {
	SavePrice(price value.Price)
}

func NewTick(tickStore repository.TickStore) Tick {
	return &tick{tickStore: tickStore}
}

// tick - ティックサービス
type tick struct {
	tickStore repository.TickStore
}

// SavePrice - 価格情報をストアに保存する
func (s *tick) SavePrice(price value.Price) {
	s.tickStore.Add(price, price.Time.Format("20060102150400"))
}
