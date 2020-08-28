package value

// SymbolLeg - 銘柄の足設定
type SymbolLeg struct {
	SymbolCode string   // 銘柄コード
	Exchange   Exchange // 市場
	LegPeriod  int      // 足の長さ(分)
}

// Is - 一致しているかどうか
func (s *SymbolLeg) Is(t SymbolLeg) bool {
	return s.SymbolCode == t.SymbolCode && s.Exchange == t.Exchange && s.LegPeriod == t.LegPeriod
}

// Exchange - 市場
type Exchange string

const (
	ExchangeUnspecified Exchange = ""  // 指定なし
	ExchangeT           Exchange = "T" // 東証
	ExchangeM           Exchange = "M" // 名証
	ExchangeF           Exchange = "F" // 福証
	ExchangeS           Exchange = "S" // 札証
)
