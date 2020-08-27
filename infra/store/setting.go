package store

import "sync"

// setting - 設定ストア
type setting struct {
	password string
	mtx      sync.Mutex
}

// IsPasswordSet - パスワードが設定済みか
func (s *setting) IsPasswordSet() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.password != ""
}
