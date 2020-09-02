package controller

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_NewPrice(t *testing.T) {
	t.Parallel()
	priceUseCase := &testPriceUseCase{}
	want := &price{
		out:          os.Stdout,
		priceUseCase: priceUseCase,
	}
	got := NewPrice(priceUseCase)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_price_Stop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		stopGetPrice error
		want         string
	}{
		{name: "エラーが返されたらエラーの内容を表示する",
			stopGetPrice: errors.New("error message"),
			want:         "価格情報の取得停止でエラーが発生しました(error message)\n"},
		{name: "エラーがなければ正常終了を表示する",
			want: "価格情報の取得を停止しました\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &price{out: w, priceUseCase: &testPriceUseCase{stopGetPrice: test.stopGetPrice}}
			controller.Stop(bufio.NewScanner(bytes.NewBufferString("")))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_price_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		startGetPrice   error
		startReturnWait time.Duration
		want            string
	}{
		{name: "エラーが返されたらエラーの内容を表示する",
			startGetPrice:   errors.New("error message"),
			startReturnWait: 100 * time.Millisecond,
			want:            "価格情報の取得を開始します\n価格情報の取得でエラーが発生しました(error message)\n"},
		{name: "エラーがなければエラー内容の表示をしない",
			want: "価格情報の取得を開始します\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &price{out: w, priceUseCase: &testPriceUseCase{startGetPrice: test.startGetPrice, startReturnWait: test.startReturnWait}}
			controller.Start(bufio.NewScanner(bytes.NewBufferString("")))
			time.Sleep(time.Second)
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
