package value

import "time"

// Price - 現値情報
type Price struct {
	Price float64   // 現値
	Time  time.Time // 現値時刻
}
