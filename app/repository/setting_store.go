package repository

// SettingStore - 設定のストア
type SettingStore interface {
	GetPassword() string
	SetPassword(password string)
	GetToken() string
	SetToken(token string)
	IsProd() bool
	SetIsProd(isProd bool)
}
