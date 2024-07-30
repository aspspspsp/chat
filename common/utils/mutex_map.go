package utils

import (
	"sync"
)

type MutexMap[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func NewMutexMap[K comparable, V any]() *MutexMap[K, V] {
	return &MutexMap[K, V]{
		data: make(map[K]V),
	}
}

func (sm *MutexMap[K, V]) Add(key K, val V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = val
}

func (sm *MutexMap[K, V]) Get(key K) (V, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	memberId, ok := sm.data[key]
	return memberId, ok
}
