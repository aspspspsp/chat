package memory

import (
	"sync"
)

type ChatRoomMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewChatroomMap() *ChatRoomMap {
	return &ChatRoomMap{
		data: make(map[string]string),
	}
}

func (sm *ChatRoomMap) Add(chatroomId string, memberId string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[chatroomId] = memberId
}

func (sm *ChatRoomMap) Get(chatroomId string) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	memberId, ok := sm.data[chatroomId]
	return memberId, ok
}
