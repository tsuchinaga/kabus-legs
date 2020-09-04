package usecase

import (
	"reflect"
	"testing"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

func Test_NewLeg(t *testing.T) {
	t.Parallel()
	symbolService := &testSymbolService{}
	legService := &testLegService{}
	clockService := &testClockService{}
	want := &leg{symbolService: symbolService, legService: legService, clockService: clockService}
	got := NewLeg(symbolService, legService, clockService)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_leg_CreateMinuteLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		want   []int
	}{
		{name: "銘柄がなければcreateもsaveも走らない",
			getAll: []value.SymbolLeg{},
			want:   []int{0, 0}},
		{name: "銘柄の数だけcreateとsaveが実行される",
			getAll: []value.SymbolLeg{
				{}, {}, {},
			},
			want: []int{3, 3}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolService := &testSymbolService{getAll: test.getAll}
			legService := &testLegService{}
			clockService := &testClockService{}
			usecase := &leg{symbolService: symbolService, legService: legService, clockService: clockService}
			usecase.CreateMinuteLeg()
			got := []int{legService.createOneMinuteLegNum, legService.saveMinuteLegNum}
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_leg_CreateMinutesLeg(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		getAll []value.SymbolLeg
		want   [3]int
	}{
		{name: "銘柄がなければ何も作られないし何も保存されない",
			getAll: []value.SymbolLeg{},
			want:   [3]int{0, 0, 0}},
		{name: "1分足の場合は作られても保存はされない",
			getAll: []value.SymbolLeg{
				{LegPeriod: 1},
				{LegPeriod: 1},
				{LegPeriod: 1},
			},
			want: [3]int{3, 0, 0}},
		{name: "1分足以外は作られて保存される",
			getAll: []value.SymbolLeg{
				{LegPeriod: 1},
				{LegPeriod: 2},
				{LegPeriod: 3},
			},
			want: [3]int{1, 2, 2}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			symbolService := &testSymbolService{getAll: test.getAll}
			legService := &testLegService{}
			clockService := &testClockService{}
			usecase := &leg{symbolService: symbolService, legService: legService, clockService: clockService}
			usecase.CreateMinutesLeg()
			got := [3]int{legService.createOneMinuteLegNum, legService.createMinutesLegNum, legService.saveMinuteLegNum}
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
