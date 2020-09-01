package repository

// PriceWebSocket - 現値WebSocketのインターフェース
type PriceWebSocket interface {
	Start() error
	Stop() error
	IsStarted() bool
}
