package value

import (
	"reflect"
	"testing"
)

func Test_SymbolLeg_Equal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    SymbolLeg
		t    SymbolLeg
		want bool
	}{
		{name: "すべての項目が一致していればtrue",
			s:    SymbolLeg{SymbolCode: "1234", Exchange: ExchangeT, LegPeriod: 3},
			t:    SymbolLeg{SymbolCode: "1234", Exchange: ExchangeT, LegPeriod: 3},
			want: true},
		{name: "一つでも項目が一致していなければfalse",
			s:    SymbolLeg{SymbolCode: "1234", Exchange: ExchangeT, LegPeriod: 3},
			t:    SymbolLeg{SymbolCode: "2345", Exchange: ExchangeT, LegPeriod: 3},
			want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.s.Equal(test.t)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
