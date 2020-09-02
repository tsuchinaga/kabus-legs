package service

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_price_SavePrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  value.Price
		want []value.Price
	}{
		{name: "repositoryのaddをたたく", arg: value.Price{SymbolCode: "1234"}, want: []value.Price{{SymbolCode: "1234"}}},
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
