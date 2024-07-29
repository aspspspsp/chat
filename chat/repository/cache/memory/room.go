package memory

import (
	"sync"
)

type RoomMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewRoomMap() *RoomMap {
	return &RoomMap{
		data: make(map[string]string),
	}
}

func (sm *RoomMap) Add(roomId string, memberId string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[roomId] = memberId
}

func (sm *RoomMap) Get(roomId string) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	memberId, ok := sm.data[roomId]
	return memberId, ok
}
