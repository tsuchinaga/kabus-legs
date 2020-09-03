package store

import (
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

// NewTick - ティックストアの生成
func NewTick() repository.TickStore {
	return &tick{store: map[value.Symbol]map[string][]value.Price{}}
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
