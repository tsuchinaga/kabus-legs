package view

import (
	"fmt"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// Setting - 設定ビューのインターフェース
type Setting interface {
	SettingStatus(status value.SettingStatus, err error) string
}

func NewSetting() Setting {
	return &setting{}
}

// setting - 設定関連のビュー
type setting struct{}

// SettingStatus - 設定状況のビューを作成
func (s *setting) SettingStatus(status value.SettingStatus, err error) string {
	if err != nil {
		return fmt.Sprintf("エラーが発生しました(%s)", err)
	}

	ps := "未設定"
	if status.IsPasswordSet {
		ps = "設定済み"
	}

	pr := "検証"
	if status.IsProd {
		pr = "本番"
	}

	return fmt.Sprintf("パスワード: %s, 環境: %s", ps, pr)
}
