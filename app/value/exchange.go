package value

// Exchange - 市場
type Exchange string

const (
	ExchangeUnspecified Exchange = ""  // 指定なし
	ExchangeT           Exchange = "T" // 東証
	ExchangeM           Exchange = "M" // 名証
	ExchangeF           Exchange = "F" // 福証
	ExchangeS           Exchange = "S" // 札証
)

// Order - 並び順
func (e Exchange) Order() int {
	switch e {
	case ExchangeT:
		return 1
	case ExchangeM:
		return 2
	case ExchangeF:
		return 3
	case ExchangeS:
		return 4
	default:
		return 99
	}
}
