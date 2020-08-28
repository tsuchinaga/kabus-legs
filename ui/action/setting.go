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

func SavePassword(bs *bufio.Scanner) gocli.AfterAction {
	fmt.Println("パスワードを入力してください: ")
	bs.Scan()
	setting := di.NewSettingUseCase()
	setting.SavePassword(bs.Text())
	fmt.Println("パスワードを設定しました")
	return gocli.AfterActionReturn
}
