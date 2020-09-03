package service

import (
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

// Clock - 時計サービスのインターフェース
type Clock interface {
	NowLabel() string
	PrevLabel(prev int) string
	NowTime() time.Time
	IsCreateLeg(legPeriod int) bool
}

// NewClock - 時計サービスの生成
func NewClock(c repository.Clock) Clock {
	return &clock{
		clock:    c,
		legStart: time.Date(0, 1, 1, 9, 0, 1, 0, time.Local),
		legEnd:   time.Date(0, 1, 1, 15, 1, 1, 0, time.Local),
	}
}

// clock - 時計サービス
type clock struct {
	clock    repository.Clock
	legStart time.Time
	legEnd   time.Time
}

// NowLabel - 現在日時のラベル
func (s *clock) NowLabel() string {
	return s.clock.Now().Format("20060102150400")
}

// PrevLabel - 現在日時から指定分前のラベル
func (s *clock) PrevLabel(prev int) string {
	return s.clock.Now().Add(-1 * time.Duration(prev) * time.Minute).Format("20060102150400")
}

// NowTime - 現在日時をとって時分秒だけを抽出した日時を返す
func (s *clock) NowTime() time.Time {
	now := s.clock.Now()
	return time.Date(0, 1, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
}

// 現在時刻が指定された足を作成するタイミングか
func (s *clock) IsCreateLeg(legPeriod int) bool {
	// 09:00:01以降、15:01:01以前でなければ作らない
	now := s.NowTime()
	if now.Before(s.legStart) || now.After(s.legEnd) {
		return false
	}
	return (now.Hour()*60+now.Minute()-9*60)%legPeriod == 0
}
