package repository

// KabuAPI - kabus apiを呼び出す処理群のインターフェース
type KabuAPI interface {
	GetToken() (string, error)
}
