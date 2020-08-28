package action

import (
	"bufio"
	"fmt"

	"gitlab.com/tsuchinaga/kabus-legs/ui/view"

	"gitlab.com/tsuchinaga/kabus-legs/di"

	"gitlab.com/tsuchinaga/gocli"
)

// GetSettingStatus -設定状況の表示
func PrintSettingStatus(_ *bufio.Scanner) gocli.AfterAction {
	setting := di.NewSettingUseCase()
	status, err := setting.GetSettingStatus()
	fmt.Println(view.SettingStatus(status, err))
	return gocli.AfterActionReturn
}

// SavePassword - パスワードの保存
func SavePassword(bs *bufio.Scanner) gocli.AfterAction {
	fmt.Println("パスワードを入力してください: ")
	bs.Scan()
	setting := di.NewSettingUseCase()
	setting.SavePassword(bs.Text())
	fmt.Println("パスワードを設定しました")
	return gocli.AfterActionReturn
}

// SetIsProd - 本番向きかの設定
func SetIsProd(bs *bufio.Scanner) gocli.AfterAction {
	fmt.Print("本番環境?(Y/N): ")
	bs.Scan()
	isProd := bs.Text()

	setting := di.NewSettingUseCase()
	if isProd == "Y" {
		setting.SetIsProd(true)
		fmt.Println("本番環境に設定しました")
	} else {
		setting.SetIsProd(false)
		fmt.Println("検証環境に設定しました")
	}
	return gocli.AfterActionReturn
}
