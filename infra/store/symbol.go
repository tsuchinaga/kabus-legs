package store

import (
	"sort"
	"sync"

	"gitlab.com/tsuchinaga/kabus-legs/app/repository"

	"gitlab.com/tsuchinaga/kabus-legs/app/value"
)

var (
	symbolStore    repository.SymbolStore
	symbolStoreMtx sync.Mutex
)

// GetSymbol - 銘柄と足の長さのストアを取得する
func GetSymbol() repository.SymbolStore {
	symbolStoreMtx.Lock()
	defer symbolStoreMtx.Unlock()

	if symbolStore == nil {
		symbolStore = &symbol{store: []value.SymbolLeg{}}
	}
	return symbolStore
}

// symbol - 銘柄と足の長さのストア
type symbol struct {
	store []value.SymbolLeg
	mtx   sync.Mutex
}

// IsExists - 引数の銘柄足が存在するか
func (s *symbol) IsExists(symbolLeg value.SymbolLeg) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, sl := range s.store {
		if sl.Is(symbolLeg) {
			return true
		}
	}
	return false
}

// Add - 銘柄足を追加する
func (s *symbol) Add(symbolLeg value.SymbolLeg) {
	if !s.IsExists(symbolLeg) {
		s.mtx.Lock()
		defer s.mtx.Unlock()
		s.store = append(s.store, symbolLeg)
		sort.Slice(s.store, func(i, j int) bool {
			return s.store[i].SymbolCode < s.store[j].SymbolCode ||
				(s.store[i].SymbolCode == s.store[j].SymbolCode && s.store[i].Exchange.Order() < s.store[j].Exchange.Order()) ||
				(s.store[i].SymbolCode == s.store[j].SymbolCode && s.store[i].Exchange.Order() == s.store[j].Exchange.Order() && s.store[i].LegPeriod < s.store[j].LegPeriod)
		})
	}
}

// GetAll - すべての銘柄足を取得する
func (s *symbol) GetAll() []value.SymbolLeg {
	return s.store
}

// DeleteByIndex - インデックスを指定して削除する
func (s *symbol) DeleteByIndex(index int) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if index < 0 || len(s.store) <= index {
		return
	}

	s.store = append(s.store[:index], s.store[index+1:]...)
}
