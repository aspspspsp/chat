package routes

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gochat/rabbitmq"
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
