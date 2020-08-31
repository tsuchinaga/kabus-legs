package service

import (
	"gitlab.com/tsuchinaga/kabus-legs/app"
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// symbol - 銘柄足サービスのインターフェース
type Symbol interface {
	GetAll() []value.SymbolLeg
	GetByIndex(index int) (value.SymbolLeg, error)
	GetBySymbol(symbolCode string, exchange value.Exchange) []value.SymbolLeg
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
func (s *symbol) GetAll() []value.SymbolLeg {
	return s.symbolStore.GetAll()
}

// AddSymbol - ストアに引数の銘柄足を追加する
func (s *symbol) AddSymbol(symbolLeg value.SymbolLeg) {
	s.symbolStore.Add(symbolLeg)
}

// DeleteSymbolByIndex - インデックスを指定して銘柄足を削除する
func (s *symbol) DeleteSymbolByIndex(index int) {
	s.symbolStore.DeleteByIndex(index)
}

// SendRegister - kabusapiを使って銘柄を登録する
func (s *symbol) SendRegister(symbolCode string, exchange value.Exchange) error {
	return s.kabuAPI.RegisterSymbol(symbolCode, exchange)
}

// SendUnregister - kabusapiに銘柄登録解除を送る
func (s *symbol) SendUnregister(symbolCode string, exchange value.Exchange) error {
	return s.kabuAPI.UnregisterSymbol(symbolCode, exchange)
}

// GetByIndex - インデックスを指定してデータを取り出す
func (s *symbol) GetByIndex(index int) (value.SymbolLeg, error) {
	symbols := s.symbolStore.GetAll()
	if index < 0 || len(symbols) <= index {
		return value.SymbolLeg{}, app.DataNotFoundError
	}
	return symbols[index], nil
}

// GetBySymbol - シンボルを指定して該当するデータをストアから取り出す
func (s *symbol) GetBySymbol(symbolCode string, exchange value.Exchange) []value.SymbolLeg {
	res := make([]value.SymbolLeg, 0)
	for _, sl := range s.symbolStore.GetAll() {
		if sl.SymbolCode == symbolCode && sl.Exchange == exchange {
			res = append(res, sl)
		}
	}
	return res
}
