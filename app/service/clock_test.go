package service

import (
	"reflect"
	"testing"
	"time"
)

func Test_NewClock(t *testing.T) {
	t.Parallel()
	c := &testClock{}
	want := &clock{clock: c}
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
