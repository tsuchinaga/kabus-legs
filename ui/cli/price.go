package cli

import (
	"bufio"

	"gitlab.com/tsuchinaga/kabus-legs/di"

	"gitlab.com/tsuchinaga/gocli"
)

// startGetPrice - 時価情報取得の開始
func startGetPrice(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewPriceController().Start(bs)
}

// stopGetPrice - 時価情報取得の停止
func stopGetPrice(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewPriceController().Stop(bs)
}
