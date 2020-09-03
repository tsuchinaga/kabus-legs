package value

import (
	"reflect"
	"testing"
	"time"
)

func Test_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{name: "2020/09/02 23:33:55をパースできる",
			time: time.Date(2020, 9, 2, 23, 33, 55, 0, time.Local),
			want: "20200902233300"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			price := &Price{Time: test.time}
			got := price.Label()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_Price_Symbol(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		price Price
		want  Symbol
	}{
		{name: "銘柄コードと市場から銘柄情報を作って取り出す",
			price: Price{SymbolCode: "1234", Exchange: ExchangeT},
			want:  Symbol{SymbolCode: "1234", Exchange: ExchangeT}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.price.Symbol()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
