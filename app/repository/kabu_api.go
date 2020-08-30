package repository

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

// KabuAPI - kabus apiを呼び出す処理群のインターフェース
type KabuAPI interface {
	GetToken() (string, error)
	RegisterSymbol(symbolCode string, exchange value.Exchange) error
	UnregisterSymbol(symbolCode string, exchange value.Exchange) error
}
