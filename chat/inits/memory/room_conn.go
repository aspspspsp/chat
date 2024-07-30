package memory

import (
	"github.com/gorilla/websocket"
	"sync"
)

var roomConnMap map[string]map[*websocket.Conn]bool
var mutex sync.Mutex

func Init() {
	roomConnMap = make(map[string]map[*websocket.Conn]bool)
}

func AddConn(roomId string, wbConn *websocket.Conn) {
	if roomConnMap == nil {
		roomConnMap = make(map[string]map[*websocket.Conn]bool)
	}

	if roomConnMap[roomId] == nil {
		roomConnMap[roomId] = make(map[*websocket.Conn]bool)
	}

	mutex.Lock()
	roomConnMap[roomId][wbConn] = true
	mutex.Unlock()
}

func GetMap() map[string]map[*websocket.Conn]bool {
	return roomConnMap
}

//func Get() {
//	roomToSocket.
//}
