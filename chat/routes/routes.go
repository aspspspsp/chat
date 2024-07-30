package routes

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
		v1.POST("sendMessage", api.SendMessageHandler())
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
