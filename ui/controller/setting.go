package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gitlab.com/tsuchinaga/gocli"
	"gitlab.com/tsuchinaga/kabus-legs/ui/view"

	"gitlab.com/tsuchinaga/kabus-legs/app/usecase"
)

// setting - 設定コントローラのインターフェース
type Setting interface {
	PrintSettingStatus(*bufio.Scanner) gocli.AfterAction
	SavePassword(*bufio.Scanner) gocli.AfterAction
	SetIsProd(*bufio.Scanner) gocli.AfterAction
}

// NewSetting - 設定コントローラの生成
func NewSetting(settingUseCase usecase.Setting, settingView view.Setting) Setting {
	return &setting{out: os.Stdout, settingUseCase: settingUseCase, settingView: settingView}
}

// setting - 設定のコントローラ
type setting struct {
	out            io.Writer
	settingUseCase usecase.Setting
	settingView    view.Setting
}

// GetSettingStatus -設定状況の表示
func (c *setting) PrintSettingStatus(_ *bufio.Scanner) gocli.AfterAction {
	status, err := c.settingUseCase.GetSettingStatus()
	_, _ = fmt.Fprintln(c.out, c.settingView.SettingStatus(status, err))
	return gocli.AfterActionReturn
}

// SavePassword - パスワードの保存
func (c *setting) SavePassword(bs *bufio.Scanner) gocli.AfterAction {
	_, _ = fmt.Fprintln(c.out, "パスワードを入力してください: ")
	bs.Scan()
	c.settingUseCase.SavePassword(bs.Text())
	_, _ = fmt.Fprintln(c.out, "パスワードを設定しました")
	return gocli.AfterActionReturn
}

// SetIsProd - 本番向きかの設定
func (c *setting) SetIsProd(bs *bufio.Scanner) gocli.AfterAction {
	_, _ = fmt.Fprintln(c.out, "本番環境?(Y/N): ")
	bs.Scan()
	isProd := bs.Text()

	if isProd == "Y" {
		c.settingUseCase.SetIsProd(true)
		_, _ = fmt.Fprintln(c.out, "本番環境に設定しました")
	} else {
		c.settingUseCase.SetIsProd(false)
		_, _ = fmt.Fprintln(c.out, "検証環境に設定しました")
	}
	return gocli.AfterActionReturn
}
