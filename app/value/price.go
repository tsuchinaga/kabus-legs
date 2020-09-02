package value

import "time"

// Price - 現値情報
type Price struct {
	SymbolCode string    // 銘柄コード
	Exchange   Exchange  // 市場
	Price      float64   // 現値
	Time       time.Time // 現値時刻
}

// Label - 現値から1分足にする場合ののラベルを取得する
func (p *Price) Label() string {
	return p.Time.Format("20060102150400")
}
