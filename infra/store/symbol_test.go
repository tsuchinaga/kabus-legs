package store

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_symbol_IsExists(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store []value.SymbolLeg
		arg   value.SymbolLeg
		want  bool
	}{
		{name: "存在すればtrue",
			store: []value.SymbolLeg{{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5}},
			arg:   value.SymbolLeg{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
			want:  true,
		},
		{name: "存在しなければfalse",
			store: []value.SymbolLeg{{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3}},
			arg:   value.SymbolLeg{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
			want:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &symbol{}
			store.mtx.Lock()
			go func() {
				defer store.mtx.Unlock()
				store.store = test.store
			}()
			got := store.IsExists(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store []value.SymbolLeg
		arg   value.SymbolLeg
		want  []value.SymbolLeg
	}{
		{name: "引数の要素をstoreに追加できる",
			store: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
			},
			arg: value.SymbolLeg{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 5},
			want: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 5},
			}},
		{name: "storeに同じ要素があったら追加されない",
			store: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
			},
			arg: value.SymbolLeg{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
			want: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
			}},
		{name: "追加時にソートがかかる",
			store: []value.SymbolLeg{
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeM, LegPeriod: 5},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
			},
			arg: value.SymbolLeg{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 1},
			want: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "1234", Exchange: value.ExchangeM, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 1},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &symbol{}
			store.mtx.Lock()
			go func() {
				defer store.mtx.Unlock()
				time.Sleep(100 * time.Millisecond)
				store.store = test.store
			}()
			store.Add(test.arg)
			got := store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store []value.SymbolLeg
		want  []value.SymbolLeg
	}{
		{name: "storeをそのまま返す",
			store: []value.SymbolLeg{{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3}},
			want:  []value.SymbolLeg{{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := (&symbol{store: test.store}).GetAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_DeleteByIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store []value.SymbolLeg
		arg   int
		want  []value.SymbolLeg
	}{
		{name: "指定したindexの要素を削除する",
			store: []value.SymbolLeg{
				{SymbolCode: "1234"},
				{SymbolCode: "2345"},
				{SymbolCode: "3456"},
			},
			arg: 1,
			want: []value.SymbolLeg{
				{SymbolCode: "1234"},
				{SymbolCode: "3456"},
			}},
		{name: "指定したindexに要素がなければ何もしない",
			store: []value.SymbolLeg{
				{SymbolCode: "1234"},
				{SymbolCode: "2345"},
				{SymbolCode: "3456"},
			},
			arg: 3,
			want: []value.SymbolLeg{
				{SymbolCode: "1234"},
				{SymbolCode: "2345"},
				{SymbolCode: "3456"},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			store := &symbol{}
			store.mtx.Lock()
			go func() {
				defer store.mtx.Unlock()
				time.Sleep(100 * time.Millisecond)
				store.store = test.store
			}()
			store.DeleteByIndex(test.arg)
			got := store.store
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
