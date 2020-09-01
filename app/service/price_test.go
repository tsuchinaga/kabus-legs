package service

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"
)

func Test_NewPriceWebSocket(t *testing.T) {
	t.Parallel()
	priceWebSocket := &testPriceWebSocket{}
	want := &price{priceWebSocket: priceWebSocket}
	got := NewPriceWebSocket(priceWebSocket)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_price_StartWebSocket(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		start error
		want  error
	}{
		{name: "Startを呼び結果を返す", start: app.WebSocketIsStoppedError, want: app.WebSocketIsStoppedError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceWebSocket := &testPriceWebSocket{start: test.start}
			service := &price{priceWebSocket: priceWebSocket}
			got := service.StartWebSocket()
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_price_StopWebSocket(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		stop error
		want error
	}{
		{name: "Stopを呼び結果を返す", stop: app.WebSocketIsStartedError, want: app.WebSocketIsStartedError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceWebSocket := &testPriceWebSocket{stop: test.stop}
			service := &price{priceWebSocket: priceWebSocket}
			got := service.StopWebSocket()
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
