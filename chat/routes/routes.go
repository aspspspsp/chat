package routes

import (
	"chat/handlers"
	"chat/rabbitmq"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

func publishHandler(ch *amqp.Channel, q amqp.Queue) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.DefaultQuery("message", "Hello World")

		err := rabbitmq.PublishMessage(ch, q, body)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to publish message: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "Message published: %s", body)
	}
}

func SetupRoutes(r *gin.Engine, ch *amqp.Channel, q amqp.Queue) {
	r.GET("/publish", publishHandler(ch, q))
}

func SetupWsRoutes() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	http.ListenAndServe(":8080", nil)
}
