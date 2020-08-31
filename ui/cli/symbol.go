package cli

import (
	"bufio"

	"gitlab.com/tsuchinaga/kabus-legs/di"

	"gitlab.com/tsuchinaga/gocli"
)

// symbolList - 銘柄足の一覧
func symbolList(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSymbolController().List(bs)
}

// symbolAdd - 銘柄足の追加
func symbolAdd(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSymbolController().Add(bs)
}

// symbolDelete - 銘柄足の削除
func symbolDelete(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSymbolController().Delete(bs)
}
