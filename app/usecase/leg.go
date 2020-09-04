package usecase

import (
	"fmt"

	"gitlab.com/tsuchinaga/kabus-legs/app/service"
	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

type Leg interface {
	CreateMinuteLeg()
	CreateMinutesLeg()
}

func NewLeg(symbolService service.Symbol, legService service.Leg, clockService service.Clock) Leg {
	return &leg{symbolService: symbolService, legService: legService, clockService: clockService}
}

type leg struct {
	symbolService service.Symbol
	legService    service.Leg
	clockService  service.Clock
}

func (u *leg) CreateMinuteLeg() {
	// 現在時刻から1分前のラベルを取り出す
	label := u.clockService.PrevLabel(1)

	// 登録銘柄一覧の取得
	for _, s := range u.symbolService.GetAll() {
		// 登録銘柄を対象に1分足作成・保存
		u.legService.SaveMinuteLeg(
			u.legService.CreateOneMinuteLeg(value.Symbol{SymbolCode: s.SymbolCode, Exchange: s.Exchange}, label),
		)
	}
}

func (u *leg) CreateMinutesLeg() {
	// 登録銘柄一覧の取得
	for _, s := range u.symbolService.GetAll() {
		if !u.clockService.IsCreateLeg(s.LegPeriod) { // 足作成時刻でないならスキップ
			continue
		}

		var leg value.FourPrice
		if s.LegPeriod == 1 { // 1分足だけは別の作り方になり、saveしてはいけない
			leg = u.legService.CreateOneMinuteLeg(value.Symbol{SymbolCode: s.SymbolCode, Exchange: s.Exchange}, u.clockService.PrevLabel(s.LegPeriod))
		} else {
			leg = u.legService.CreateMinutesLeg(value.Symbol{SymbolCode: s.SymbolCode, Exchange: s.Exchange}, u.clockService.PrevLabel(s.LegPeriod), s.LegPeriod)
			u.legService.SaveMinuteLeg(leg)
		}
		fmt.Printf("%+v\n", leg)
	}
}
