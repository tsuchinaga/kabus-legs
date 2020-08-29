package kabu

import "gitlab.com/tsuchinaga/go-kabusapi/kabus"

// TokenRequester - トークンリクエスタ
type TokenRequester interface {
	Exec(request kabus.TokenRequest) (*kabus.TokenResponse, error)
}

// RegisterRequester - 銘柄登録リクエスタ
type RegisterRequester interface {
	Exec(request kabus.RegisterRequest) (*kabus.RegisterResponse, error)
}

// UnregisterRequester - 銘柄登録解除リクエスタ
type UnregisterRequester interface {
	Exec(request kabus.UnregisterRequest) (*kabus.UnregisterResponse, error)
}
