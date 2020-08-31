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
