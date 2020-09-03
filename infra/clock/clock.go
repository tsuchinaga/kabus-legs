package clock

import (
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"
)

// NewClock  - 時計の生成
func NewClock() repository.Clock {
	return &clock{}
}

// clock - 時計
type clock struct{}

// Now - 現在日時
func (c *clock) Now() time.Time {
	return time.Now()
}
