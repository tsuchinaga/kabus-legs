package repository

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

// LegStore - 足ストア
type LegStore interface {
	Add(fourPrice value.FourPrice)
	Get(symbol value.Symbol, legPeriod int) []value.FourPrice
}
