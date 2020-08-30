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

func (s *symbolLeg) Unregister(int) error { panic("implement me") }
