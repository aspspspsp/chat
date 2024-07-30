package routes

import (
	api "chat/api/v1"
	"chat/api/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.POST("/sendMessage", api.SendMessageHandler())
	}
	//
	//r.POST("/addUserToChatRoom", controllers.AddMemberToChatroom)
	//r.POST("/removeUserFromChatRoom", controllers.RemoveUserFromChatRoom)

	return r
}

func SetupWsRoutes() {
	http.HandleFunc("/ws", ws.HandleWebSocket)
}
