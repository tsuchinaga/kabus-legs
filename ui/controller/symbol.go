package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"

	"gitlab.com/tsuchinaga/gocli"
)

// Symbol - 銘柄コントローラのインターフェース
type Symbol interface {
	List(*bufio.Scanner) gocli.AfterAction
	Add(*bufio.Scanner) gocli.AfterAction
	Delete(*bufio.Scanner) gocli.AfterAction
}

// NewSymbol - 新しい銘柄コントローラの生成
func NewSymbol(symbolLegUseCase usecase.SymbolLeg) Symbol {
	return &symbol{out: os.Stdout, symbolLegUseCase: symbolLegUseCase}
}

// symbol - 銘柄コントローラ
type symbol struct {
	out              io.Writer
	symbolLegUseCase usecase.SymbolLeg
}

// List - 銘柄ストアのデータを取り出して表示する
func (c *symbol) List(*bufio.Scanner) gocli.AfterAction {
	symbols, err := c.symbolLegUseCase.GetAll()
	if err != nil {
		_, _ = fmt.Fprintf(c.out, "銘柄一覧取得に失敗しました(%s)\n", err)
		return gocli.AfterActionReturn
	}
	if len(symbols) == 0 {
		_, _ = fmt.Fprintln(c.out, "登録されている銘柄と足はありません")
		return gocli.AfterActionReturn
	}

	_, _ = fmt.Fprintf(c.out, "No. | 銘柄コード | 市場 | 足の長さ\n")
	for i, s := range symbols {
		_, _ = fmt.Fprintf(c.out, "%3d | %10s | %4s | %8d\n", i, s.SymbolCode, s.Exchange, s.LegPeriod)
	}
	return gocli.AfterActionReturn
}
func (c *symbol) Add(*bufio.Scanner) gocli.AfterAction    { panic("implement me") }
func (c *symbol) Delete(*bufio.Scanner) gocli.AfterAction { panic("implement me") }
