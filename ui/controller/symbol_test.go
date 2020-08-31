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

func Test_symbol_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		register error
		in       io.Reader
		want     string
	}{
		{name: "市場コードに指定文字以外を入れるとエラー",
			in: bytes.NewBufferString("1234\nA\n"),
			want: "銘柄コードを入力してください: " +
				"市場コードを入力してください(T: 東証, M: 名証, F: 福証, S: 札証): " +
				"市場はT, M, F, Sで入力してください\n"},
		{name: "足の長さに数字以外を入れるとエラー",
			in: bytes.NewBufferString("1234\nT\n５\n"),
			want: "銘柄コードを入力してください: " +
				"市場コードを入力してください(T: 東証, M: 名証, F: 福証, S: 札証): " +
				"足の長さを入力してください(分): " +
				"足の長さは半角数字で入力してください\n"},
		{name: "registerに失敗するとエラー",
			in:       bytes.NewBufferString("1234\nT\n5\n"),
			register: errors.New("error message"),
			want: "銘柄コードを入力してください: " +
				"市場コードを入力してください(T: 東証, M: 名証, F: 福証, S: 札証): " +
				"足の長さを入力してください(分): " +
				"銘柄登録でエラーが発生しました(error message)\n"},
		{name: "銘柄登録に成功すると成功メッセージ",
			in: bytes.NewBufferString("1234\nT\n5\n"),
			want: "銘柄コードを入力してください: " +
				"市場コードを入力してください(T: 東証, M: 名証, F: 福証, S: 札証): " +
				"足の長さを入力してください(分): " +
				"銘柄登録に成功しました\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &symbol{out: w, symbolLegUseCase: &testSymbolLegUseCase{register: test.register}}
			controller.Add(bufio.NewScanner(test.in))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
