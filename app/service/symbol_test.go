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

func Test_symbol_SendUnregister(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		unregisterSymbol error
		arg1             string
		arg2             value.Exchange
		want             error
	}{
		{name: "repositoryからエラーが返されたらそのまま返す",
			unregisterSymbol: app.APIRequestError,
			want:             app.APIRequestError},
		{name: "repositoryからエラーがなくてもそのまま返す"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			kabuAPI := &testKabusAPI{unregisterSymbol: test.unregisterSymbol}
			service := &symbol{kabuAPI: kabuAPI}
			got := service.SendUnregister(test.arg1, test.arg2)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbol_GetByIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		arg    int
		want1  value.SymbolLeg
		want2  error
	}{
		{name: "指定したインデックスが存在しなければerror",
			getAll: []value.SymbolLeg{{SymbolCode: "1234"}},
			arg:    1,
			want2:  app.DataNotFoundError},
		{name: "指定したインデックスが存在すれば該当のvalueを返す",
			getAll: []value.SymbolLeg{{SymbolCode: "1234"}, {SymbolCode: "5678"}},
			arg:    1,
			want1:  value.SymbolLeg{SymbolCode: "5678"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolStore := &testSymbolStore{getAll: test.getAll}
			service := &symbol{symbolStore: symbolStore}
			got1, got2 := service.GetByIndex(test.arg)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_symbol_GetBySymbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		arg1   string
		arg2   value.Exchange
		want   []value.SymbolLeg
	}{
		{name: "該当するデータがなければ空スライス",
			getAll: []value.SymbolLeg{},
			arg1:   "1234", arg2: value.ExchangeT,
			want: []value.SymbolLeg{}},
		{name: "該当するデータがあればスライスに入れて返す",
			getAll: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 1},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeM, LegPeriod: 1},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 1},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 10},
			},
			arg1: "1234", arg2: value.ExchangeT,
			want: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 1},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolStore := &testSymbolStore{getAll: test.getAll}
			service := &symbol{symbolStore: symbolStore}
			got := service.GetBySymbol(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
