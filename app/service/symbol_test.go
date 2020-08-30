package service

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_Name(t *testing.T) {
	t.Parallel()
	symbolStore := &testSymbolStore{}
	kabuAPI := &testKabusAPI{}
	want := &symbol{
		symbolStore: symbolStore,
		kabuAPI:     kabuAPI,
	}
	got := NewSymbol(symbolStore, kabuAPI)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_symbol_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		want   []value.SymbolLeg
	}{
		{name: "storeが返した結果をそのまま返す"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolStore := &testSymbolStore{getAll: test.getAll}
			service := &symbol{symbolStore: symbolStore}
			got := service.GetAll()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_AddSymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  value.SymbolLeg
		want []value.SymbolLeg
	}{
		{name: "銘柄足ストアに追加する",
			arg:  value.SymbolLeg{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 1},
			want: []value.SymbolLeg{{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 1}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolStore := &testSymbolStore{}
			service := &symbol{symbolStore: symbolStore}
			service.AddSymbol(test.arg)
			got := symbolStore.addHis
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_DeleteSymbolByIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  int
		want []int
	}{
		{name: "インデックス番号をストアに渡す", arg: 3, want: []int{3}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolStore := &testSymbolStore{}
			service := &symbol{symbolStore: symbolStore}
			service.DeleteSymbolByIndex(test.arg)
			got := symbolStore.deleteByIndexHis
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_RegisterSymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		registerSymbol error
		arg1           string
		arg2           value.Exchange
		want           error
	}{
		{name: "repositoryからエラーが返されたらそのまま返す",
			registerSymbol: app.APIRequestError,
			want:           app.APIRequestError},
		{name: "repositoryからエラーがなくてもそのまま返す"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			kabuAPI := &testKabusAPI{registerSymbol: test.registerSymbol}
			service := &symbol{kabuAPI: kabuAPI}
			got := service.SendRegister(test.arg1, test.arg2)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
