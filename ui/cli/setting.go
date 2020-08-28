package cli

import (
	"bufio"

	"gitlab.com/tsuchinaga/gocli"
	"gitlab.com/tsuchinaga/kabus-legs/di"
)

// printSettingStatus - 設定状況の表示
func printSettingStatus(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSettingController().PrintSettingStatus(bs)
}

// savePassword - パスワードの保存
func savePassword(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSettingController().SavePassword(bs)
}

// setIsProd - 本番環境か検証環境かの保存
func setIsProd(bs *bufio.Scanner) gocli.AfterAction {
	return di.NewSettingController().SetIsProd(bs)
}
