package store

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_tick_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store map[value.Symbol]map[string][]value.Price
		arg1  value.Price
		arg2  string
		want  map[value.Symbol]map[string][]value.Price
	}{
		{name: "未初期化のストアに価格情報を突っ込んだら新たにMapが作られて追加される",
			store: nil,
			arg1: value.Price{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			},
			arg2: "20200903092200",
			want: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {"20200903092200": []value.Price{{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			}}}}},
		{name: "銘柄が初めての場合でも銘柄のMapを作って追加される",
			store: map[value.Symbol]map[string][]value.Price{},
			arg1: value.Price{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			},
			arg2: "20200903092200",
			want: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {"20200903092200": []value.Price{{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			}}}}},
		{name: "ラベルが初めての場合でもラベルのMapを作って追加される",
			store: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {}},
			arg1: value.Price{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			},
			arg2: "20200903092200",
			want: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {"20200903092200": []value.Price{{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			}}}}},
		{name: "既存のMapがあればそこに追加される",
			store: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {"20200903092100": []value.Price{}}},
			arg1: value.Price{
				SymbolCode: "1234",
				Exchange:   value.ExchangeT,
				Price:      23000,
				Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
			},
			arg2: "20200903092200",
			want: map[value.Symbol]map[string][]value.Price{value.Symbol{SymbolCode: "1234", Exchange: value.ExchangeT}: {
				"20200903092100": []value.Price{},
				"20200903092200": []value.Price{{
					SymbolCode: "1234",
					Exchange:   value.ExchangeT,
					Price:      23000,
					Time:       time.Date(2020, 9, 3, 9, 22, 17, 0, time.Local),
				}}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &tick{store: test.store}
			store.Add(test.arg1, test.arg2)
			got := store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_NewTick(t *testing.T) {
	t.Parallel()
	want := &tick{store: map[value.Symbol]map[string][]value.Price{}}
	got := NewTick()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
