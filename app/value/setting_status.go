package value

// SettingStatus - 設定状況
type SettingStatus struct {
	IsPasswordSet bool // パスワードが設定されているか
	IsProd        bool // 本番を向いているか
}
