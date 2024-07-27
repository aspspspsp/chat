package main

import (
	"chat/config"
	"chat/rabbitmq"
	"chat/routes"
	"common/utils"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	// db初始化
	config.DbInit()

	// mq初始化
	conn, ch, q := rabbitmq.ConnectRabbitMQ()
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			utils.FailOnError(err, "chat服務啟動失敗")
		}
	}(conn)

	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			utils.FailOnError(err, "member服務啟動失敗")
		}
	}(ch)

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
		utils.FailOnError(err, "Failed to register a consumer")

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
	err := r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
