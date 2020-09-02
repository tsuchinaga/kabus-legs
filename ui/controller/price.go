package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"

	"gitlab.com/tsuchinaga/gocli"
)

// Price - 時価情報コントローラのインターフェース
type Price interface {
	Start(*bufio.Scanner) gocli.AfterAction
	Stop(*bufio.Scanner) gocli.AfterAction
}

// NewPrice - 時価情報コントローラの生成
func NewPrice(priceUseCase usecase.Price) Price {
	return &price{out: os.Stdout, priceUseCase: priceUseCase}
}

// price - 時価情報コントローラ
type price struct {
	out          io.Writer
	priceUseCase usecase.Price
}

// Start - 時価情報の取得開始
func (c *price) Start(_ *bufio.Scanner) gocli.AfterAction {
	go func() {
		if err := c.priceUseCase.StartGetPrice(); err != nil {
			_, _ = fmt.Fprintf(c.out, "価格情報の取得でエラーが発生しました(%s)\n", err)
		}
	}()
	_, _ = fmt.Fprintln(c.out, "価格情報の取得を開始します")
	return gocli.AfterActionReturn
}

// Start - 時価情報の取得停止
func (c *price) Stop(_ *bufio.Scanner) gocli.AfterAction {
	if err := c.priceUseCase.StopGetPrice(); err != nil {
		_, _ = fmt.Fprintf(c.out, "価格情報の取得停止でエラーが発生しました(%s)\n", err)
	} else {
		_, _ = fmt.Fprintln(c.out, "価格情報の取得を停止しました")
	}
	return gocli.AfterActionReturn
}
