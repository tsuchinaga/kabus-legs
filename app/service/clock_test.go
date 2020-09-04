package service

import (
	"reflect"
	"testing"
	"time"
)

func Test_NewClock(t *testing.T) {
	t.Parallel()
	c := &testClock{}
	want := &clock{
		clock:    c,
		legStart: time.Date(0, 1, 1, 9, 0, 1, 0, time.Local),
		legEnd:   time.Date(0, 1, 1, 15, 1, 1, 0, time.Local),
	}
	got := NewClock(c)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_clock_NowLabel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		now  time.Time
		want string
	}{
		{name: "現在時刻のラベルが取れる",
			now:  time.Date(2020, 9, 3, 14, 53, 22, 0, time.Local),
			want: "20200903145300"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &clock{clock: &testClock{now: test.now}}
			got := service.NowLabel()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_clock_NowTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		now  time.Time
		want time.Time
	}{
		{name: "日時から時分秒を取り出してtime.Timeを作る",
			now:  time.Date(2020, 9, 3, 14, 53, 22, 0, time.Local),
			want: time.Date(0, 1, 1, 14, 53, 22, 0, time.Local)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			service := &clock{clock: &testClock{now: test.now}}
			got := service.NowTime()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_clock_IsCreateLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		now  time.Time
		arg  int
		want bool
	}{
		{name: "09:00:00はfalse",
			now:  time.Date(0, 1, 1, 9, 0, 0, 0, time.Local),
			arg:  1,
			want: false},
		{name: "15:02:00はfalse",
			now:  time.Date(0, 1, 1, 15, 2, 0, 0, time.Local),
			arg:  1,
			want: false},
		{name: "09:03:00に3分足はtrue",
			now:  time.Date(0, 1, 1, 9, 3, 0, 0, time.Local),
			arg:  3,
			want: true},
		{name: "09:03:00に5分足はfalse",
			now:  time.Date(0, 1, 1, 9, 3, 0, 0, time.Local),
			arg:  5,
			want: false},
		{name: "14:50:00に5分足はtrue",
			now:  time.Date(0, 1, 1, 14, 50, 0, 0, time.Local),
			arg:  5,
			want: true},
		{name: "14:50:00に20分足はfalse",
			now:  time.Date(0, 1, 1, 14, 50, 0, 0, time.Local),
			arg:  20,
			want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewClock(&testClock{now: test.now}).IsCreateLeg(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_clock_PrevLabel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		now  time.Time
		arg  int
		want string
	}{
		{name: "09:22:00の10分前なら09:12:00のラベルが出る",
			now:  time.Date(2020, 9, 3, 9, 22, 0, 0, time.Local),
			arg:  10,
			want: "20200903091200"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := (&clock{clock: &testClock{now: test.now}}).PrevLabel(test.arg)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
