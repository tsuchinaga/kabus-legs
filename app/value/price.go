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
// 本来現値の責務ではないから使わない方向で
func (p *Price) Label() string {
	return p.Time.Format("20060102150400")
}

// Symbol - 現値情報から銘柄情報を取り出す
func (p *Price) Symbol() Symbol {
	return Symbol{
		SymbolCode: p.SymbolCode,
		Exchange:   p.Exchange,
	}
}
