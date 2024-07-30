package message_broadcast

import (
	"chat/repository/db/models"
	"chat/repository/mq/message_store"
	"common/repository/mq"
	"common/utils"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const messageExchangeName = "messageBroadcasting"

var amqpChannel *amqp.Channel

func InitMq() {
	conn := mq.ConnectRabbitMQ()
	ch := mq.CreateChannel(conn)
	defer conn.Close()
	defer ch.Close()

	declareExchange(ch)
	q := declareQueue(ch)
	mq.BindQueue(ch, q.Name, "", messageExchangeName)
	msgs := consumeMessages(ch, q)

	go handleMessages(msgs)
}

func declareExchange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		messageExchangeName, // 交换机名字
		"fanout",            // 交换机类型，这里使用fanout类型，即: 发布订阅模式
		true,                // 是否持久化
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	utils.FailOnError(err, "Failed to declare an exchange")
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")
	return q
}

// PublishMessage 發送聊天消息
func PublishMessage(ch *amqp.Channel, message models.Message) {
	body, err := json.Marshal(message)

	err = ch.Publish(
		messageExchangeName, // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func consumeMessages(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
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
	return msgs
}

func handleMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		body := d.Body
		log.Printf("Received a message: %s", body)
		var message models.Message
		err := json.Unmarshal(body, &message)
		if err != nil {
			log.Fatalf("Error decoding JSON: %s", err)
			continue // 跳过当前消息，继续处理下一个消息
		}

		// 消息解藕
		message_store.PublishMessage(amqpChannel, message)
	}
}
