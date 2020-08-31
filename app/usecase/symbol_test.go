package usecase

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_symbolLeg_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		want1  []value.SymbolLeg
		want2  error
	}{
		{name: "serviceが返した結果をそのまま返す",
			getAll: []value.SymbolLeg{{SymbolCode: "1234"}, {SymbolCode: "5678"}, {SymbolCode: "9012"}},
			want1:  []value.SymbolLeg{{SymbolCode: "1234"}, {SymbolCode: "5678"}, {SymbolCode: "9012"}},
			want2:  nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolService := &testSymbolService{getAll: test.getAll}
			usecase := &symbolLeg{symbolService: symbolService}
			got1, got2 := usecase.GetAll()
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_symbolLeg_Register(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sendRegister error
		arg1         string
		arg2         string
		arg3         int
		want         error
	}{
		{name: "sendRegisterに失敗したらエラーが返される"},
		{name: "sendRegisterに成功したらストアに登録して正常終了"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolService := &testSymbolService{sendRegister: test.sendRegister}
			usecase := &symbolLeg{symbolService: symbolService}
			got := usecase.Register(test.arg1, test.arg2, test.arg3)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbolLeg_Unregister(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		getByIndex1    value.SymbolLeg
		getByIndex2    error
		getBySymbol    []value.SymbolLeg
		sendUnregister error
		arg            int
		want           error
	}{
		{name: "該当インデックスが存在しなければエラー",
			getByIndex2: app.DataNotFoundError,
			want:        app.DataNotFoundError},
		{name: "銘柄登録解除APIをたたいてエラーが返されればエラー",
			getBySymbol:    []value.SymbolLeg{{SymbolCode: "1234"}},
			sendUnregister: app.APIRequestError,
			want:           app.APIRequestError},
		{name: "同一銘柄が2つ以上あればunregisterはたたかない",
			getBySymbol:    []value.SymbolLeg{{SymbolCode: "1234"}, {SymbolCode: "1234"}},
			sendUnregister: app.APIRequestError,
			want:           nil},
		{name: "該当インデックスが存在して登録解除に成功したらストアからも消して正常終了",
			want: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolService := &testSymbolService{getByIndex1: test.getByIndex1, getByIndex2: test.getByIndex2, getBySymbol: test.getBySymbol, sendUnregister: test.sendUnregister}
			usecase := &symbolLeg{symbolService: symbolService}
			got := usecase.Unregister(test.arg)
			if !errors.Is(got, test.want) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_NewSymbolLeg(t *testing.T) {
	t.Parallel()
	symbolService := &testSymbolService{}
	want := &symbolLeg{symbolService: symbolService}
	got := NewSymbolLeg(symbolService)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
