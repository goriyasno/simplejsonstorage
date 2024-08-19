package simplejsonstorage

import (
	"encoding/json"
	"os"
	"sync"
)

type Storage[T any] struct {
	Lock sync.RWMutex
	Path string
}

func New[T any](path string) (s *Storage[T]) {
	return &Storage[T]{
		Path: path,
	}
}

func (s *Storage[T]) Read(p *T) (err error) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	f, err := os.ReadFile(s.Path)
	if err != nil {
		return
	}
	return json.Unmarshal(f, p)
}

func (s *Storage[T]) Write(p *T) (err error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	return os.WriteFile(s.Path, data, os.ModePerm)
}
