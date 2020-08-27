package kabu

import "gitlab.com/tsuchinaga/go-kabusapi/kabus"

// TokenRequester - トークンリクエスタ
type TokenRequester interface {
	Exec(request kabus.TokenRequest) (*kabus.TokenResponse, error)
}
