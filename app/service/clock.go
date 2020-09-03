package service

import "gitlab.com/tsuchinaga/kabus-legs/app/repository"

// Clock - 時計サービスのインターフェース
type Clock interface {
	NowLabel() string
}

// NewClock - 時計サービスの生成
func NewClock(c repository.Clock) Clock {
	return &clock{clock: c}
}

// clock - 時計サービス
type clock struct {
	clock repository.Clock
}

// NowLabel - 現在日時のラベル
func (s *clock) NowLabel() string {
	return s.clock.Now().Format("20060102150400")
}
