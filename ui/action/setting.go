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
