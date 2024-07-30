package mq

import (
	"common/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() *amqp.Connection {
	// TODO: 使用連線池
	// RabbitMQ 连接，使用正确的用户名和密码
	conn, err := amqp.Dial("amqp://twg:123456@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	return ch
}

func BindQueue(ch *amqp.Channel, queueName string, routingKey string, exchangeName string) {
	err := ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")
}
