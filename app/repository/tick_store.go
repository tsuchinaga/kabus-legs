package repository

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

// TickStore - 価格ストアのインターフェース
type TickStore interface {
	Add(price value.Price, label string)
	Get(symbol value.Symbol, label string) []value.Price
}
