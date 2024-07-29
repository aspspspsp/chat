package routes

import (
	"chat/api/ws"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

//func publishHandler(ch *amqp.Channel, q amqp.Queue) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		body := c.DefaultQuery("message", "Hello World")
//
//		err := mq.PublishMessage(ch, q, body)
//		if err != nil {
//			c.String(http.StatusInternalServerError, "Failed to publish message: %s", err.Error())
//			return
//		}
//
//		c.String(http.StatusOK, "Message published: %s", body)
//	}
//}

func SetupRoutes(r *gin.Engine, ch *amqp.Channel, q amqp.Queue) {
	//r.GET("/publish", publishHandler(ch, q))
	//
	//r.POST("/addUserToChatRoom", controllers.AddMemberToChatroom)
	//r.POST("/removeUserFromChatRoom", controllers.RemoveUserFromChatRoom)
}

func SetupWsRoutes() {
	http.HandleFunc("/ws", ws.HandleWebSocket)
	http.ListenAndServe(":8080", nil)
}
