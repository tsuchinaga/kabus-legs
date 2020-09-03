package store

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

var (
	legStore    repository.LegStore
	legStoreMtx sync.Mutex
)

// GetLeg - 足ストアを取得する
func GetLeg() repository.LegStore {
	legStoreMtx.Lock()
	defer legStoreMtx.Unlock()

	if legStore == nil {
		legStore = &leg{store: map[value.Symbol]map[int][]value.FourPrice{}}
	}
	return legStore
}

type leg struct {
	store map[value.Symbol]map[int][]value.FourPrice
	mtx   sync.Mutex
}

// Add - 足情報をストアに登録する
func (s *leg) Add(fourPrice value.FourPrice) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.store == nil {
		s.store = map[value.Symbol]map[int][]value.FourPrice{}
	}

	if _, ok := s.store[fourPrice.Symbol]; !ok {
		s.store[fourPrice.Symbol] = map[int][]value.FourPrice{}
	}

	if _, ok := s.store[fourPrice.Symbol][fourPrice.LegPeriod]; !ok {
		s.store[fourPrice.Symbol][fourPrice.LegPeriod] = []value.FourPrice{}
	}

	s.store[fourPrice.Symbol][fourPrice.LegPeriod] = append(s.store[fourPrice.Symbol][fourPrice.LegPeriod], fourPrice)
}

// Get - 指定した四本値のsliceを取り出す
func (s *leg) Get(symbol value.Symbol, legPeriod int) []value.FourPrice {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.store == nil {
		return []value.FourPrice{}
	}

	if _, ok := s.store[symbol]; !ok {
		return []value.FourPrice{}
	}

	if _, ok := s.store[symbol][legPeriod]; !ok {
		return []value.FourPrice{}
	}

	return s.store[symbol][legPeriod]
}
