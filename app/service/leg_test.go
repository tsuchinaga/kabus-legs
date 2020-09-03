package service

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_leg_CreateOneMinuteLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		tickGet []value.Price
		legGet1 []value.FourPrice
		arg1    value.Symbol
		arg2    string
		want    value.FourPrice
	}{
		{name: "tick情報がなければ1本前の足から現値をとってきて仮想足をつくる",
			tickGet: []value.Price{},
			legGet1: []value.FourPrice{{Close: 20000}},
			arg1:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:    "20200903091500",
			want: value.FourPrice{
				Symbol:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Label:     "20200903091500",
				Open:      20000,
				High:      20000,
				Low:       20000,
				Close:     20000,
				LegPeriod: 1,
				IsVirtual: true,
			}},
		{name: "tick情報も1本前の足もなければ0円で仮想足を作る",
			tickGet: []value.Price{},
			legGet1: []value.FourPrice{},
			arg1:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:    "20200903091500",
			want: value.FourPrice{
				Symbol:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Label:     "20200903091500",
				Open:      0,
				High:      0,
				Low:       0,
				Close:     0,
				LegPeriod: 1,
				IsVirtual: true,
			}},
		{name: "tick情報があれば最初の足を始値、最後の足を終値とし、後は高値と安値を設定して足を作る",
			tickGet: []value.Price{
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20000},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20005},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20010},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20005},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20000},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 19995},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 19990},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20000},
				{SymbolCode: "1234", Exchange: value.ExchangeT, Price: 20005},
			},
			legGet1: []value.FourPrice{},
			arg1:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:    "20200903091500",
			want: value.FourPrice{
				Symbol:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Label:     "20200903091500",
				Open:      20000,
				High:      20010,
				Low:       19990,
				Close:     20005,
				LegPeriod: 1,
				IsVirtual: false,
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			tickStore := &testTickStore{get: test.tickGet}
			legStore := &testLegStore{get1: test.legGet1}
			service := &leg{tickStore: tickStore, legStore: legStore}
			got := service.CreateOneMinuteLeg(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_leg_SaveMinuteLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  value.FourPrice
		want []value.FourPrice
	}{
		{name: "引数をaddに渡す",
			arg: value.FourPrice{
				Symbol:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Label:     "20200903142500",
				Open:      20000,
				High:      20500,
				Low:       20000,
				Close:     20250,
				LegPeriod: 1,
				IsVirtual: false,
			},
			want: []value.FourPrice{{
				Symbol:    value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Label:     "20200903142500",
				Open:      20000,
				High:      20500,
				Low:       20000,
				Close:     20250,
				LegPeriod: 1,
				IsVirtual: false,
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			legStore := &testLegStore{}
			service := &leg{legStore: legStore}
			service.SaveMinuteLeg(test.arg)
			got := legStore.addHis
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_leg_CreateMinutesLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		get1 []value.FourPrice
		getN []value.FourPrice
		arg1 value.Symbol
		arg2 string
		arg3 int
		want value.FourPrice
	}{
		{name: "1分足の本数が足りず、同じ長さの足も存在しない場合、0の値を持つ四本値が返される",
			get1: []value.FourPrice{{}, {}},
			getN: []value.FourPrice{},
			arg1: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2: "20200903143900",
			arg3: 3,
			want: value.FourPrice{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, Label: "20200903143900", LegPeriod: 3, IsVirtual: true}},
		{name: "1分足の本数が足りない場合、同じ長さの足をとってきて、存在したら終値を使って仮想足を作る",
			get1: []value.FourPrice{{}, {}},
			getN: []value.FourPrice{{Close: 22000}},
			arg1: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2: "20200903143900",
			arg3: 3,
			want: value.FourPrice{
				Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Open:   22000, High: 22000, Low: 22000, Close: 22000,
				Label: "20200903143900", LegPeriod: 3, IsVirtual: true}},
		{name: "1分足が必要な本数あれば、そこから四本値を作る",
			get1: []value.FourPrice{
				{Open: 22000, High: 22005, Low: 21990, Close: 21990},
				{Open: 21990, High: 21990, Low: 21970, Close: 21975},
				{Open: 21970, High: 22000, Low: 21965, Close: 21990}},
			arg1: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2: "20200903143900",
			arg3: 3,
			want: value.FourPrice{
				Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Open:   22000, High: 22005, Low: 21965, Close: 21990,
				Label: "20200903143900", LegPeriod: 3, IsVirtual: false}},
		{name: "1分足が必要な本数あれば、そこから四本値を作るけど始値0はスキップされる",
			get1: []value.FourPrice{
				{},
				{Open: 21990, High: 21990, Low: 21970, Close: 21975},
				{Open: 21970, High: 22000, Low: 21965, Close: 21990}},
			arg1: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2: "20200903143900",
			arg3: 3,
			want: value.FourPrice{
				Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
				Open:   21990, High: 22000, Low: 21965, Close: 21990,
				Label: "20200903143900", LegPeriod: 3, IsVirtual: false}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			legStore := &testLegStore{get1: test.get1, getN: test.getN}
			service := &leg{legStore: legStore}
			got := service.CreateMinutesLeg(test.arg1, test.arg2, test.arg3)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_NewLeg(t *testing.T) {
	t.Parallel()
	tickStore := &testTickStore{}
	legStore := &testLegStore{}
	want := &leg{tickStore: tickStore, legStore: legStore}
	got := NewLeg(tickStore, legStore)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
