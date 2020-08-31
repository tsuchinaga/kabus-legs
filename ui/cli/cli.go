package cli

import (
	"gitlab.com/tsuchinaga/gocli"
)

func Run() error {
	gocli.SetCommandNotExistsMessage("コマンドがありません")
	gocli.SetHelpDescription("ヘルプメッセージ")
	gocli.SetReturnDescription("前のメニューに戻る")
	gocli.SetExitDescription("終了")
	g := gocli.NewGocli()

	// 設定
	g.AddSubCommand(gocli.NewCommand("setting", "設定").
		AddSubCommand(gocli.NewCommand("status", "設定状況").SetAction(printSettingStatus)).
		AddSubCommand(gocli.NewCommand("password", "パスワードの設定").SetAction(savePassword)).
		AddSubCommand(gocli.NewCommand("prod", "本番向きかの設定").SetAction(setIsProd)),
	)

	// トークン
	g.AddSubCommand(gocli.NewCommand("token", "トークン").
		AddSubCommand(gocli.NewCommand("get", "現在のトークンの取得").SetAction(getToken)).
		AddSubCommand(gocli.NewCommand("new", "新しいトークンの発行").SetAction(refreshToken)).
		AddSubCommand(gocli.NewCommand("set", "任意のトークンの利用").SetAction(setToken)),
	)

	// 銘柄足
	g.AddSubCommand(gocli.NewCommand("symbol", "銘柄").
		AddSubCommand(gocli.NewCommand("list", "登録している銘柄の一覧").SetAction(symbolList)).
		AddSubCommand(gocli.NewCommand("add", "足作成する銘柄の追加").SetAction(symbolAdd)).
		AddSubCommand(gocli.NewCommand("delete", "足作成する銘柄の削除").SetAction(symbolDelete)))

	return g.Run()
}
