package value

import (
	"reflect"
	"testing"
)

func Test_Exchange_Order(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		exchange Exchange
		want     int
	}{
		{name: "Tが1", exchange: ExchangeT, want: 1},
		{name: "Mが2", exchange: ExchangeM, want: 2},
		{name: "Fが3", exchange: ExchangeF, want: 3},
		{name: "Sが4", exchange: ExchangeS, want: 4},
		{name: "空文字が99", exchange: ExchangeUnspecified, want: 99},
		{name: "任意の文字列が99", exchange: Exchange("foo"), want: 99},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.exchange.Order()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
