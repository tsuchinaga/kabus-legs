package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"

	"gitlab.com/tsuchinaga/gocli"
)

// Token - トークンコントローラのインターフェース
type Token interface {
	GetToken(_ *bufio.Scanner) gocli.AfterAction
	RefreshToken(_ *bufio.Scanner) gocli.AfterAction
	SetToken(bs *bufio.Scanner) gocli.AfterAction
}

// NewToken - トークンコントローラの生成
func NewToken(settingUsecase usecase.Setting) Token {
	return &token{out: os.Stdout, settingUseCase: settingUsecase}
}

// token - トークンコントローラ
type token struct {
	out            io.Writer
	settingUseCase usecase.Setting
}

// GetToken - トークンの取得
func (c *token) GetToken(_ *bufio.Scanner) gocli.AfterAction {
	token, err := c.settingUseCase.GetToken()
	if err != nil {
		_, _ = fmt.Fprintf(c.out, "エラーが発生しました(%s)\n", err)
		return gocli.AfterActionReturn
	}
	if token == "" {
		_, _ = fmt.Fprintln(c.out, "トークン未取得")
	} else {
		_, _ = fmt.Fprintln(c.out, token)
	}
	return gocli.AfterActionReturn
}

// RefreshToken - トークンの再取得
func (c *token) RefreshToken(_ *bufio.Scanner) gocli.AfterAction {
	token, err := c.settingUseCase.GetNewToken()
	if err != nil {
		_, _ = fmt.Fprintf(c.out, "エラーが発生しました(%s)\n", err)
		return gocli.AfterActionReturn
	}
	if token == "" {
		_, _ = fmt.Fprintln(c.out, "トークン再発行に失敗しました")
	} else {
		_, _ = fmt.Fprintln(c.out, token)
	}
	return gocli.AfterActionReturn
}

// SetToken - トークンの設定
func (c *token) SetToken(bs *bufio.Scanner) gocli.AfterAction {
	_, _ = fmt.Fprintln(c.out, "セットするトークンを入力してください: ")
	bs.Scan()
	token := bs.Text()
	c.settingUseCase.SaveToken(token)
	_, _ = fmt.Fprintln(c.out, "トークンをセットしました")
	return gocli.AfterActionReturn
}
