package value

// SymbolLeg - 銘柄の足設定
type SymbolLeg struct {
	SymbolCode string   // 銘柄コード
	Exchange   Exchange // 市場
	LegPeriod  int      // 足の長さ(分)
}

// Equal - 一致しているかどうか
func (s *SymbolLeg) Equal(t SymbolLeg) bool {
	return s.SymbolCode == t.SymbolCode && s.Exchange == t.Exchange && s.LegPeriod == t.LegPeriod
}
