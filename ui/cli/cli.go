package cli

import (
	"gitlab.com/tsuchinaga/gocli"
	"gitlab.com/tsuchinaga/kabus-legs/ui/action"
)

func Run() error {
	g := gocli.NewGocli()

	// 設定
	g.AddSubCommand(gocli.NewCommand("setting", "設定").
		AddSubCommand(gocli.NewCommand("status", "設定状況").SetAction(action.PrintSettingStatus)).
		AddSubCommand(gocli.NewCommand("password", "パスワードの設定")).
		AddSubCommand(gocli.NewCommand("prod", "本番向きかの設定")),
	)

	// トークン
	g.AddSubCommand(gocli.NewCommand("token", "トークン").
		AddSubCommand(gocli.NewCommand("get", "現在のトークンの取得")).
		AddSubCommand(gocli.NewCommand("new", "新しいトークンの発行")).
		AddSubCommand(gocli.NewCommand("set", "任意のトークンの利用")),
	)

	return g.Run()
}
