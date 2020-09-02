package value

// FourPrice - 四本値
type FourPrice struct {
	Symbol    Symbol  // 銘柄
	Label     string  // ラベル
	Open      float64 // 始値
	High      float64 // 高値
	Low       float64 // 安値
	Close     float64 // 終値
	IsVirtual bool    // 仮想足
}
