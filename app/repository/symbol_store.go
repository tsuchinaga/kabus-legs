package repository

import "gitlab.com/tsuchinaga/kabus-legs/app/value"

// SymbolStore - 銘柄と足の長さのストアのインターフェース
type SymbolStore interface {
	IsExists(symbolLeg value.SymbolLeg) bool
	Add(symbolLeg value.SymbolLeg)
	GetAll() []value.SymbolLeg
	DeleteByIndex(index int)
}
