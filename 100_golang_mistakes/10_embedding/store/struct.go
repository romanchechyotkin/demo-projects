package store

import "sync"

// store mutex functions like Lock() or Unlock() would be public via embedding
type store struct {
	store map[string]string
	sync.Mutex
}

func NewStore() *store {
	return &store{store: make(map[string]string)}
}

func (s *store) Get() {}

// store2 mutex functions like Lock() or Unlock() would be private via non export field
type store2 struct {
	store map[string]string
	mu    sync.Mutex
}

func NewStore2() *store2 {
	return &store2{store: make(map[string]string), mu: sync.Mutex{}}
}

func (s *store2) Get() {}
