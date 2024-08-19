package simplejsonstorage

import (
	"encoding/json"
	"os"
	"sync"
)

type Storage struct {
	Lock sync.RWMutex
	Path string
}

func New(path string) (s *Storage) {
	return &Storage{
		Path: path,
	}
}

func (s *Storage) Read(p any) (err error) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	f, err := os.ReadFile(s.Path)
	if err != nil {
		return
	}
	return json.Unmarshal(f, p)
}

func (s *Storage) Write(p any) (err error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	return os.WriteFile(s.Path, data, os.ModePerm)
}
