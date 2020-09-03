package clock

import (
	"reflect"
	"testing"
	"time"
)

func Test_NewClock(t *testing.T) {
	t.Parallel()
	want := &clock{}
	got := NewClock()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_clock_Now(t *testing.T) {
	t.Parallel()
	want := time.Now()
	c := &clock{}
	got := c.Now()
	if want.After(got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
