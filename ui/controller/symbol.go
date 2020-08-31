package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

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

// Add - 銘柄登録
func (c *symbol) Add(bs *bufio.Scanner) gocli.AfterAction {
	_, _ = fmt.Fprint(c.out, "銘柄コードを入力してください: ")
	bs.Scan()
	code := bs.Text()

	_, _ = fmt.Fprint(c.out, "市場コードを入力してください(T: 東証, M: 名証, F: 福証, S: 札証): ")
	bs.Scan()
	exchange := bs.Text()
	if exchange != "T" && exchange != "M" && exchange != "F" && exchange != "S" {
		_, _ = fmt.Fprintln(c.out, "市場はT, M, F, Sで入力してください")
		return gocli.AfterActionReturn
	}

	_, _ = fmt.Fprint(c.out, "足の長さを入力してください(分): ")
	bs.Scan()
	legPeriod, err := strconv.Atoi(bs.Text())
	if err != nil {
		_, _ = fmt.Fprintln(c.out, "足の長さは半角数字で入力してください")
		return gocli.AfterActionReturn
	}

	if err := c.symbolLegUseCase.Register(code, exchange, legPeriod); err != nil {
		_, _ = fmt.Fprintf(c.out, "銘柄登録でエラーが発生しました(%s)\n", err)
		return gocli.AfterActionReturn
	}

	_, _ = fmt.Fprintln(c.out, "銘柄登録に成功しました")
	return gocli.AfterActionReturn
}
func (c *symbol) Delete(*bufio.Scanner) gocli.AfterAction { panic("implement me") }
