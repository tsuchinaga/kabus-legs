package usecase

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// SymbolLeg - 銘柄足ユースケースのインターフェース
type SymbolLeg interface {
	GetAll() ([]value.SymbolLeg, error)
	Register(symbolCode string, exchange string, legPeriod int) error
	Unregister(index int) error
}

// NewSymbolLeg - 銘柄足ユースケースの生成
func NewSymbolLeg(symbolService service.Symbol) SymbolLeg {
	return &symbolLeg{symbolService: symbolService}
}

// symbolLeg - 銘柄足ユースケース
type symbolLeg struct {
	symbolService service.Symbol
}

// GetAll - ストアに登録された銘柄一足覧を取り出す
func (s *symbolLeg) GetAll() ([]value.SymbolLeg, error) {
	return s.symbolService.GetAll(), nil
}

// Register - 銘柄足を作成し、ストアへの登録とAPIへの登録を行う
func (s *symbolLeg) Register(symbolCode string, exchange string, legPeriod int) error {
	if err := s.symbolService.SendRegister(symbolCode, value.Exchange(exchange)); err != nil {
		return err
	}

	s.symbolService.AddSymbol(value.SymbolLeg{SymbolCode: symbolCode, Exchange: value.Exchange(exchange), LegPeriod: legPeriod})
	return nil
}

// Unregister - インデックスを指定して銘柄の登録を解除する
func (s *symbolLeg) Unregister(index int) error {
	symbol, err := s.symbolService.GetByIndex(index)
	if err != nil {
		return err
	}

	if err := s.symbolService.SendUnregister(symbol.SymbolCode, symbol.Exchange); err != nil {
		return err
	}

	s.symbolService.DeleteSymbolByIndex(index)
	return nil
}
