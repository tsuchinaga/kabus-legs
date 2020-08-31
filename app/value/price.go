package value

import "time"

// Price - 現値情報
type Price struct {
	SymbolCode string    // 銘柄コード
	Exchange   Exchange  // 市場
	Price      float64   // 現値
	Time       time.Time // 現値時刻
}
