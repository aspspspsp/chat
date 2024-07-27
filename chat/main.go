package main

import (
	"chat/rabbitmq"
	"chat/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, ch, q := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Gin 路由设置
	r := gin.Default()

	routes.SetupRoutes(r, ch, q)

	routes.SetupWsRoutes()

	go func() {
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}()

	// 启动 HTTP 服务
	r.Run(":8080")
}
