package service

import (
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// leg - 足サービスのインターフェース
type Leg interface {
	CreateOneMinuteLeg(symbol value.Symbol, label string) value.FourPrice
	CreateMinutesLeg(symbol value.Symbol, label string, legPeriod int) value.FourPrice
	SaveMinuteLeg(fourPrice value.FourPrice)
}

// leg - 足サービス
type leg struct {
	tickStore repository.TickStore
	legStore  repository.LegStore
}

// CreateOneMinuteLeg - 現値情報から1分足を生成する
func (s *leg) CreateOneMinuteLeg(symbol value.Symbol, label string) value.FourPrice {
	ticks := s.tickStore.Get(symbol, label)
	if len(ticks) == 0 {
		legs := s.legStore.Get(symbol, 1)
		if len(legs) == 0 {
			return value.FourPrice{
				Symbol: symbol, Label: label, Open: 0, High: 0, Low: 0, Close: 0, LegPeriod: 1, IsVirtual: true,
			}
		} else {
			leg := legs[len(legs)-1]
			return value.FourPrice{
				Symbol: symbol, Label: label, Open: leg.Close, High: leg.Close, Low: leg.Close, Close: leg.Close, LegPeriod: 1, IsVirtual: true,
			}
		}
	}

	four := value.FourPrice{
		Symbol: symbol, Label: label, Open: ticks[0].Price, High: ticks[0].Price, Low: ticks[0].Price, Close: ticks[len(ticks)-1].Price, LegPeriod: 1, IsVirtual: false,
	}
	for _, t := range ticks {
		if four.High < t.Price {
			four.High = t.Price
		}
		if four.Low > t.Price {
			four.Low = t.Price
		}
	}
	return four
}

// CreateMinutesLeg - 任意の長さの足を1分足から生成する
func (s *leg) CreateMinutesLeg(symbol value.Symbol, label string, legPeriod int) value.FourPrice {
	legs := s.legStore.Get(symbol, 1)
	if len(legs) < legPeriod {
		legs := s.legStore.Get(symbol, legPeriod)
		if len(legs) == 0 {
			return value.FourPrice{
				Symbol: symbol, Label: label, Open: 0, High: 0, Low: 0, Close: 0, LegPeriod: legPeriod, IsVirtual: true,
			}
		} else {
			leg := legs[len(legs)-1]
			return value.FourPrice{
				Symbol: symbol, Label: label, Open: leg.Close, High: leg.Close, Low: leg.Close, Close: leg.Close, LegPeriod: legPeriod, IsVirtual: true,
			}
		}
	}

	four := value.FourPrice{
		Symbol:    symbol,
		Label:     label,
		LegPeriod: legPeriod,
	}
	for i := 0; i < legPeriod; i++ {
		leg := legs[len(legs)-legPeriod+i]
		if leg.Open == 0 { // 始値0はゴミデータなのでスキップする
			continue
		}

		if four.Open == 0 {
			four.Open = leg.Open
		}
		if four.High < leg.High || four.High == 0 {
			four.High = leg.High
		}
		if four.Low > leg.Low || four.Low == 0 {
			four.Low = leg.Low
		}
		four.Close = leg.Close
	}
	return four
}

// SaveMinuteLeg - 四本値情報を保存する
func (s *leg) SaveMinuteLeg(fourPrice value.FourPrice) {
	s.legStore.Add(fourPrice)
}
