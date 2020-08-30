package service

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// symbol - 銘柄足サービスのインターフェース
type Symbol interface {
	GetAll() []value.SymbolLeg
	AddSymbol(symbolLeg value.SymbolLeg)
	DeleteSymbolByIndex(index int)
	SendRegister(symbolCode string, exchange value.Exchange) error
	SendUnregister(symbolCode string, exchange value.Exchange) error
}

// NewSymbol - 銘柄足サービスの生成
func NewSymbol(symbolStore repository.SymbolStore, kabuAPI repository.KabuAPI) Symbol {
	return &symbol{symbolStore: symbolStore, kabuAPI: kabuAPI}
}

// symbol - 銘柄足サービス
type symbol struct {
	symbolStore repository.SymbolStore
	kabuAPI     repository.KabuAPI
}

// GetAll - ストアに保持されているデータを取得する
func (s symbol) GetAll() []value.SymbolLeg {
	return s.symbolStore.GetAll()
}

// AddSymbol - ストアに引数の銘柄足を追加する
func (s symbol) AddSymbol(symbolLeg value.SymbolLeg) {
	s.symbolStore.Add(symbolLeg)
}

func (s symbol) DeleteSymbolByIndex(int)                     { panic("implement me") }
func (s symbol) SendRegister(string, value.Exchange) error   { panic("implement me") }
func (s symbol) SendUnregister(string, value.Exchange) error { panic("implement me") }
