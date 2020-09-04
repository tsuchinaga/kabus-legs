package scheduler

import (
	"sync"
	"time"

	"gitlab.com/tsuchinaga/kabus-legs/infra/clock"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/di"
)

func Run() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			t := NewCreateLeg()
			n := t.Next()
			<-time.After(n)
			t.Run()
		}
	}()

	wg.Wait()
}

type Task interface {
	Run()
	Next() time.Duration
}

func NewCreateLeg() Task {
	return &createLeg{clock: clock.NewClock()}
}

// createLeg - 足作成処理
// 毎分3秒に実行
// 09:01:00 ~ 15:01:00 の間動作する。ただし11:30~12:30の間は昼休憩なので動かない
// 現在時刻が09:01:03以前なら次は09:01:03に、期間中なら次の分の03秒に、15:01:00以降なら動かず終わり
type createLeg struct {
	clock repository.Clock
}

func (t *createLeg) Run() {
	usecase := di.NewLegUseCase()
	usecase.CreateMinuteLeg()  // 1分足作成
	usecase.CreateMinutesLeg() // N分足作成
}

func (t *createLeg) Next() time.Duration {
	now := t.clock.Now()
	if now.Before(time.Date(now.Year(), now.Month(), now.Day(), 9, 1, 3, 0, time.Local)) {
		return time.Date(now.Year(), now.Month(), now.Day(), 9, 1, 3, 0, time.Local).Sub(now)
	} else if now.After(time.Date(now.Year(), now.Month(), now.Day(), 11, 31, 0, 0, time.Local)) &&
		now.Before(time.Date(now.Year(), now.Month(), now.Day(), 12, 31, 3, 0, time.Local)) {
		return time.Date(now.Year(), now.Month(), now.Day(), 12, 31, 3, 0, time.Local).Sub(now)
	} else if now.After(time.Date(now.Year(), now.Month(), now.Day(), 15, 1, 0, 0, time.Local)) {
		next := now.AddDate(0, 0, 1)
		return time.Date(next.Year(), next.Month(), next.Day(), 9, 1, 3, 0, time.Local).Sub(now)
	}

	next := now.Add(1 * time.Minute)
	return time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 3, 0, time.Local).Sub(now)
}
