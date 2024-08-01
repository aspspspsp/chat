package ws

import (
	"chat/inits/memory"
	"common/repository/db/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有连接
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	//defer conn.Close()

	// 从查询参数中获取聊天室名称
	_roomId := r.URL.Query().Get("roomId")
	if _roomId == "" {
		log.Println("No room specified")
	}

	// 将 roomId 转换为整数
	roomId, err := strconv.Atoi(_roomId)
	if err != nil {
		http.Error(w, "Invalid roomId", http.StatusBadRequest)
		return
	}

	err = conn.WriteMessage(1, []byte("歡迎光臨聊天室"))

	if err != nil {
		panic(err)
		return
	}

	memory.AddConn(uint(roomId), conn)

	for {
		// 读取消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading msg:", err)
			break
		}
		log.Printf("Received: %s", msg)

		// 回传消息
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println("Error writing msg:", err)
			break
		}
	}
}

func BroadcastMessage(message models.Message) {
	roomId := message.RoomId
	log.Println("roomId:" + strconv.Itoa(int(roomId)))
	body, _ := json.Marshal(message)

	for c := range memory.GetMap()[roomId] {

		if err := c.WriteMessage(1, []byte(body)); err != nil {
			log.Println("Error writing message to user:", err)
			//c.Close()
			//delete(chatRooms[room], c)
			// 如果聊天室为空，则删除该聊天室
			//if len(chatRooms[room]) == 0 {
			//	delete(chatRooms, room)
			//}
		}
	}
}
