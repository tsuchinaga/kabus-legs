package service

import (
	"reflect"
	"testing"

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
