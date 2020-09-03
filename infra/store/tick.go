package store

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

var (
	tickStore    repository.TickStore
	tickStoreMtx sync.Mutex
)

// GetTick - ティックストアの取得
func GetTick() repository.TickStore {
	tickStoreMtx.Lock()
	defer tickStoreMtx.Unlock()
	if tickStore == nil {
		tickStore = &tick{store: map[value.Symbol]map[string][]value.Price{}}
	}
	return tickStore
}

// tick - ティックストア
type tick struct {
	store map[value.Symbol]map[string][]value.Price
	mtx   sync.Mutex
}

// Add - 現値情報を追加する
func (s *tick) Add(price value.Price, label string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.store == nil {
		s.store = map[value.Symbol]map[string][]value.Price{}
	}

	if _, ok := s.store[price.Symbol()]; !ok {
		s.store[price.Symbol()] = map[string][]value.Price{}
	}

	if _, ok := s.store[price.Symbol()][label]; !ok {
		s.store[price.Symbol()][label] = []value.Price{}
	}

	s.store[price.Symbol()][label] = append(s.store[price.Symbol()][label], price)
}

// Get - 指定した現値情報のsliceを取得する
func (s *tick) Get(symbol value.Symbol, label string) []value.Price {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.store == nil {
		return []value.Price{}
	}

	if _, ok := s.store[symbol]; !ok {
		return []value.Price{}
	}

	if _, ok := s.store[symbol][label]; !ok {
		return []value.Price{}
	}

	return s.store[symbol][label]
}
