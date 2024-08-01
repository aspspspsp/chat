package route

import (
	api "chat/api/v1"
	"chat/api/ws"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		chat := v1.Group("message")
		{
			chat.POST("send", api.SendMessageHandler())
		}
		room := v1.Group("room")
		{
			room.POST("create", api.CreateHandler())
			room.POST("addMember", api.AddToRoomHandler())
			room.DELETE("removeMember", api.RemoveToRoomHandler())
		}

		v1.GET("test", api.TestHandler())
	}
	//
	//r.POST("/addUserToChatRoom", controllers.AddMemberToChatroom)
	//r.POST("/removeUserFromChatRoom", controllers.RemoveUserFromChatRoom)

	return r
}

func SetupWsRoutes() {
	http.HandleFunc("/ws", ws.HandleWebSocket)

	go func() {
		err := http.ListenAndServe(":7777", nil)
		if err != nil {
			log.Println("websocket 啟動失敗")
		}
	}()
}
