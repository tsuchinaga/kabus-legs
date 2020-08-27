package kabu

import (
	"fmt"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"gitlab.com/tsuchinaga/kabus-legs/app"
	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

// kabu - kabus apiを呼び出す処理をまとめた構造体
type kabu struct {
	settingStore   repository.SettingStore
	tokenRequester TokenRequester
}

// GetToken - トークンの取得
func (k *kabu) GetToken() (string, error) {
	res, err := k.tokenRequester.Exec(kabus.TokenRequest{APIPassword: k.settingStore.GetPassword()})
	if err != nil {
		return "", fmt.Errorf("%v: %w", err, app.APIRequestError)
	}
	return res.Token, nil
}
