package usecase

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"
)

func Test_NewPrice(t *testing.T) {
	priceService := &testPriceService{}
	want := &price{priceService: priceService}
	got := NewPrice(priceService)
	t.Parallel()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_price_StartGetPrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		startWebSocket error
		want           error
	}{
		{name: "serviceをたたきその結果を返す", startWebSocket: app.WebSocketIsStartedError, want: app.WebSocketIsStartedError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceService := &testPriceService{startWebSocket: test.startWebSocket}
			usecase := &price{priceService: priceService}
			got := usecase.StartGetPrice()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_price_StopGetPrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		stopWebSocket error
		want          error
	}{
		{name: "serviceをたたきその結果を返す", stopWebSocket: app.WebSocketIsStoppedError, want: app.WebSocketIsStoppedError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceService := &testPriceService{stopWebSocket: test.stopWebSocket}
			usecase := &price{priceService: priceService}
			got := usecase.StopGetPrice()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
