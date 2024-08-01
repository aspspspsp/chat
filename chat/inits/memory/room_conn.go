package memory

import (
	"github.com/gorilla/websocket"
	"sync"
)

var roomConnMap map[uint]map[*websocket.Conn]bool
var mutex sync.Mutex

func AddConn(roomId uint, wbConn *websocket.Conn) {
	if roomConnMap == nil {
		roomConnMap = make(map[uint]map[*websocket.Conn]bool)
	}

	if roomConnMap[roomId] == nil {
		roomConnMap[roomId] = make(map[*websocket.Conn]bool)
	}

	mutex.Lock()
	roomConnMap[roomId][wbConn] = true
	mutex.Unlock()
}

func GetMap() map[uint]map[*websocket.Conn]bool {
	return roomConnMap
}

//func Get() {
//	roomToSocket.
//}
