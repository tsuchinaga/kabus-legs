package service

import (
	"reflect"
	"testing"
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
