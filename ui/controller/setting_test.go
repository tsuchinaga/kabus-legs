package controller

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_NewSetting(t *testing.T) {
	t.Parallel()
	settingUseCase := &testSettingUseCase{}
	settingView := &testSettingView{}
	want := &setting{
		out:            os.Stdout,
		settingUseCase: settingUseCase,
		settingView:    settingView,
	}
	got := NewSetting(settingUseCase, settingView)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_setting_PrintSettingStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		settingStatus string
		want          string
	}{
		{name: "viewが返した結果をそのまま出力に渡す",
			settingStatus: "view message",
			want:          "view message\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &setting{out: w, settingUseCase: &testSettingUseCase{}, settingView: &testSettingView{settingStatus: test.settingStatus}}
			controller.PrintSettingStatus(bufio.NewScanner(os.Stdin))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SavePassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   io.Reader
		want string
	}{
		{name: "パスワード保存処理が実行されて表示される",
			in:   bytes.NewBufferString("password\n"),
			want: "パスワードを入力してください: \nパスワードを設定しました\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &setting{out: w, settingUseCase: &testSettingUseCase{}}
			controller.SavePassword(bufio.NewScanner(test.in))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_setting_SetIsProd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   io.Reader
		want string
	}{
		{name: "入力が検証環境の場合、検証環境として登録される",
			in:   bytes.NewBufferString("N\n"),
			want: "本番環境?(Y/N): \n検証環境に設定しました\n"},
		{name: "入力が本番環境の場合、本番環境として登録される",
			in:   bytes.NewBufferString("Y\n"),
			want: "本番環境?(Y/N): \n本番環境に設定しました\n"},
		{name: "入力がYでもNでもない場合、検証環境として登録される",
			in:   bytes.NewBufferString("foo\n"),
			want: "本番環境?(Y/N): \n検証環境に設定しました\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := new(bytes.Buffer)
			controller := &setting{out: w, settingUseCase: &testSettingUseCase{}}
			controller.SetIsProd(bufio.NewScanner(test.in))
			got := w.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
