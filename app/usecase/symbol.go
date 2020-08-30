package usecase

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// SymbolLeg - 銘柄足ユースケースのインターフェース
type SymbolLeg interface {
	GetAll() ([]value.SymbolLeg, error)
	Add(symbolCode string, exchange value.Exchange, legPeriod int) error
	Delete(index int) error
}

// symbolLeg - 銘柄足ユースケース
type symbolLeg struct {
	symbolService service.Symbol
}

// GetAll - ストアに登録された銘柄一足覧を取り出す
func (s *symbolLeg) GetAll() ([]value.SymbolLeg, error) {
	return s.symbolService.GetAll(), nil
}

func (s *symbolLeg) Add(string, value.Exchange, int) error { panic("implement me") }
func (s *symbolLeg) Delete(int) error                      { panic("implement me") }
