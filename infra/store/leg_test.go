package store

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_leg_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store map[value.Symbol]map[int][]value.FourPrice
		arg   value.FourPrice
		want  map[value.Symbol]map[int][]value.FourPrice
	}{
		{name: "storeがnilなら作って放り込む",
			store: nil,
			arg:   value.FourPrice{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1},
			want:  map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {1: []value.FourPrice{{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1}}}}},
		{name: "storeにsymbolがなければ作って放り込む",
			store: map[value.Symbol]map[int][]value.FourPrice{},
			arg:   value.FourPrice{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1},
			want:  map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {1: []value.FourPrice{{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1}}}}},
		{name: "storeにlegPeriodがなければ作って放り込む",
			store: map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {}},
			arg:   value.FourPrice{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1},
			want:  map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {1: []value.FourPrice{{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1}}}}},
		{name: "storeがあれば末尾に追加する",
			store: map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {1: []value.FourPrice{{}}}},
			arg:   value.FourPrice{Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1},
			want:  map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {1: []value.FourPrice{{}, {Symbol: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}, LegPeriod: 1}}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &leg{store: test.store}
			store.Add(test.arg)
			got := store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_leg_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store map[value.Symbol]map[int][]value.FourPrice
		arg1  value.Symbol
		arg2  int
		want  []value.FourPrice
	}{
		{name: "storeがなければ空sliceが返される",
			store: nil,
			arg1:  value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:  1,
			want:  []value.FourPrice{}},
		{name: "銘柄がなければ空sliceが返される",
			store: map[value.Symbol]map[int][]value.FourPrice{},
			arg1:  value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:  1,
			want:  []value.FourPrice{}},
		{name: "足の長さがなければ空sliceが返される",
			store: map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {}},
			arg1:  value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2:  1,
			want:  []value.FourPrice{}},
		{name: "あればsliceが返される",
			store: map[value.Symbol]map[int][]value.FourPrice{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {
				1: []value.FourPrice{{Label: "20200903090000"}, {Label: "20200903090100"}, {Label: "20200903090200"}},
			}},
			arg1: value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT},
			arg2: 1,
			want: []value.FourPrice{{Label: "20200903090000"}, {Label: "20200903090100"}, {Label: "20200903090200"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &leg{store: test.store}
			got := store.Get(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_GetLeg(t *testing.T) {
	t.Parallel()
	want := &leg{store: map[value.Symbol]map[int][]value.FourPrice{}}
	got := GetLeg()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
