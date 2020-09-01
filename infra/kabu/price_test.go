package kabu

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/app"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_NewPrice(t *testing.T) {
	// t.Parallel() // グローバル変数にアクセスするので直列テスト
	tests := []struct {
		name    string
		priceWS repository.PriceWebSocket
		want1   bool
		want2   bool
	}{
		{name: "priceWSがnilなら新たに生成されて返される",
			priceWS: nil,
			want1:   true},
		{name: "生成済みならWebSocketIsStartedErrorが返される",
			priceWS: &price{},
			want2:   true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel() // グローバル変数にアクセスするので直列テスト
			priceWS = test.priceWS
			got1, got2 := NewPrice(&testSettingStore{}, func(v value.Price) error { return nil })
			if test.want1 != (got1 != nil) || test.want2 != (got2 != nil) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_GetPrice(t *testing.T) {
	// t.Parallel() // グローバル変数にアクセスするので直列テスト
	tests := []struct {
		name    string
		priceWS repository.PriceWebSocket
		want1   bool
		want2   bool
	}{
		{name: "priceWSがあればそれを返す",
			priceWS: &price{},
			want1:   true},
		{name: "priceWSがなければエラー",
			priceWS: nil,
			want2:   true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel() // グローバル変数にアクセスするので直列テスト
			priceWS = test.priceWS
			got1, got2 := GetPrice()
			if test.want1 != (got1 != nil) || test.want2 != (got2 != nil) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_price_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		started bool
		open    error
		want    error
	}{
		{name: "すでに開始されている場合はエラー", started: true, want: app.WebSocketIsStartedError},
		{name: "requesterが返した結果を返す", open: errors.New("error message"), want: app.APIRequestError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceWSRequester := &testPriceWSRequester{open: test.open}
			ws := &price{priceWSRequester: priceWSRequester, started: test.started}
			got := ws.Start()
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_price_Stop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		close   error
		started bool
		want    error
	}{
		{name: "すでに止まっていたらエラー", want: app.WebSocketIsStoppedError},
		{name: "requesterの返した結果を返す", started: true, close: errors.New("error message"), want: app.APIRequestError},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			priceWSRequester := &testPriceWSRequester{close: test.close}
			ws := &price{priceWSRequester: priceWSRequester, started: test.started}
			got := ws.Stop()
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_price_IsStarted(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		started bool
		want    bool
	}{
		{name: "startedをそのままかえす", started: true, want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ws := &price{started: test.started}
			got := ws.IsStarted()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
