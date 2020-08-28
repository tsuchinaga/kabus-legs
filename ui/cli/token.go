package cli

import (
	"bufio"

	"gitlab.com/tsuchinaga/kabus-legs/di"

	"gitlab.com/tsuchinaga/gocli"
)

// getToken - トークンの取得
func getToken(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewTokenController().GetToken(bs)
}

// refreshToken - トークンの再取得
func refreshToken(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewTokenController().RefreshToken(bs)
}

// setToken - トークンのセット
func setToken(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewTokenController().SetToken(bs)
}
