package memory

import (
	"sync"
)

type ChatroomMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewChatroomMap() *ChatroomMap {
	return &ChatroomMap{
		data: make(map[string]string),
	}
}

func (sm *ChatroomMap) Add(chatroomId string, memberId string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[chatroomId] = memberId
}

func (sm *ChatroomMap) Get(chatroomId string) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	memberId, ok := sm.data[chatroomId]
	return memberId, ok
}
