package controller

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_NewToken(t *testing.T) {
	t.Parallel()
	testSettingUseCase := &testSettingUseCase{}
	want := &token{out: os.Stdout, settingUseCase: testSettingUseCase}
	got := NewToken(testSettingUseCase)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_setting_GetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		getToken1 string
		getToken2 error
		want      string
	}{
		{name: "エラーが返されたらそのエラーを表示する",
			getToken2: errors.New("error message"),
			want:      "エラーが発生しました(error message)\n"},
		{name: "トークンが空文字なら未取得扱いする",
			getToken1: "",
			want:      "トークン未取得\n"},
		{name: "トークンに有効な文字列が入っていればトークンを表示する",
			getToken1: "token",
			want:      "token\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &token{out: w, settingUseCase: &testSettingUseCase{getToken1: test.getToken1, getToken2: test.getToken2}}
			controller.GetToken(bufio.NewScanner(os.Stdin))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_RefreshToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		getNewToken1 string
		getNewToken2 error
		want         string
	}{
		{name: "エラーが発生したらエラーを表示する",
			getNewToken2: errors.New("error message"),
			want:         "エラーが発生しました(error message)\n"},
		{name: "トークンが取れたらトークンを表示する",
			getNewToken1: "token",
			want:         "token\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &token{out: w, settingUseCase: &testSettingUseCase{getNewToken1: test.getNewToken1, getNewToken2: test.getNewToken2}}
			controller.RefreshToken(bufio.NewScanner(os.Stdin))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_token_SetToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   io.Reader
		want string
	}{
		{name: "入力されたトークンを受け取る",
			in:   bytes.NewBufferString("token\n"),
			want: "セットするトークンを入力してください: \nトークンをセットしました\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &token{out: w, settingUseCase: &testSettingUseCase{}}
			controller.SetToken(bufio.NewScanner(test.in))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
