package store

import (
	"sync"
	"time"
)

type Store struct {
	mu     sync.RWMutex
	data   map[string]string
	expiry map[string]time.Time
}

func NewStore() *Store {
	return &Store{
		data:   make(map[string]string),
		expiry: make(map[string]time.Time),
	}
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value

}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok

}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *Store) Expiry(key string, seconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.expiry[key] = time.Now().Add(time.Duration(seconds) * time.Second)
}
func (s *Store) StartSweeper() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			s.mu.Lock()
			for key, expireAt := range s.expiry {
				if time.Now().After(expireAt) {
					delete(s.data, key)
					delete(s.expiry, key)
				}

			}
			s.mu.Unlock()

		}
	}()
}
