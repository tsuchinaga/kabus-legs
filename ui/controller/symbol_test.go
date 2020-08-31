package controller

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_symbol_List(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		getAll1 []value.SymbolLeg
		getAll2 error
		in      io.Reader
		want    string
	}{
		{name: "errorが返されたらエラーの内容を表示",
			getAll2: errors.New("error message"),
			want:    "銘柄一覧取得に失敗しました(error message)\n"},
		{name: "銘柄が登録されていなければ登録なしを表示する",
			getAll1: []value.SymbolLeg{},
			want:    "登録されている銘柄と足はありません\n"},
		{name: "銘柄が返されたらいい感じにスペースでそろえて表示する",
			getAll1: []value.SymbolLeg{
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "1234", Exchange: value.ExchangeT, LegPeriod: 5},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 3},
				{SymbolCode: "5678", Exchange: value.ExchangeT, LegPeriod: 10},
			},
			want: `No. | 銘柄コード | 市場 | 足の長さ
  0 |       1234 |    T |        3
  1 |       1234 |    T |        5
  2 |       5678 |    T |        3
  3 |       5678 |    T |       10
`},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &symbol{out: w, symbolLegUseCase: &testSymbolLegUseCase{getAll1: test.getAll1, getAll2: test.getAll2}}
			controller.List(bufio.NewScanner(test.in))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_NewSymbol(t *testing.T) {
	t.Parallel()
	symbolLegUseCase := &testSymbolLegUseCase{}
	want := &symbol{out: os.Stdout, symbolLegUseCase: symbolLegUseCase}
	got := NewSymbol(symbolLegUseCase)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
