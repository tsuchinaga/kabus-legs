package service

import (
	"reflect"
	"testing"
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_price_SavePrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  value.Price
		want []struct {
			price value.Price
			label string
		}
	}{
		{name: "repositoryのaddをたたく", arg: value.Price{SymbolCode: "1234", Time: time.Date(2020, 9, 3, 9, 15, 22, 0, time.Local)}, want: []struct {
			price value.Price
			label string
		}{{price: value.Price{SymbolCode: "1234", Time: time.Date(2020, 9, 3, 9, 15, 22, 0, time.Local)}, label: "20200903091500"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			tickStore := &testTickStore{}
			service := &tick{tickStore: tickStore}
			service.SavePrice(test.arg)
			got := tickStore.addHis
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_NewTick(t *testing.T) {
	t.Parallel()
	tickStore := &testTickStore{}
	want := &tick{tickStore: tickStore}
	got := NewTick(tickStore)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
