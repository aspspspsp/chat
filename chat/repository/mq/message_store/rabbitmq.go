package message_store

import (
	"chat/repository/db/dao"
	"common/repository/db/models"
	"common/repository/mq"
	"common/utils"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const queueName = "messageStore"

var messageDao *dao.MessageDao

func InitMq(ctx context.Context) {
	// 创建 MessageDao 实例
	messageDao = dao.NewMessageDao(ctx)

	conn := mq.ConnectRabbitMQ()
	ch := mq.CreateChannel(conn)
	// TODO: 移到main方法
	//defer conn.Close()
	//defer ch.Close()

	declareQueue(ch)

	msgs := consumeMessages(ch)
	done := make(chan bool)

	go handleMessages(ctx, msgs, done)
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")
	return q
}

func PublishMessage(ch *amqp.Channel, message models.Message) {
	body, err := json.Marshal(message)

	err = ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent(非廣播) %s", message)
}

func consumeMessages(ch *amqp.Channel) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	utils.FailOnError(err, "Failed to register a consumer")
	return msgs
}

func handleMessages(ctx context.Context, msgs <-chan amqp.Delivery, done chan bool) {
	for d := range msgs {
		body := d.Body
		log.Printf("Received a message(非廣播): %s", body)
		var message models.Message
		err := json.Unmarshal(body, &message)
		if err != nil {
			log.Printf("Error decoding JSON: %s", err)
			continue // 跳过当前消息，继续处理下一个消息
		}
		err = messageDao.Create(&message)
		if err != nil {
			log.Printf("Error 存儲失敗: %s", err)
			err := d.Nack(false, true)
			if err != nil {
				continue
			} // 发送 Nack 并重新入队
			continue // 跳过当前消息，继续处理下一个消息
		}
		d.Ack(false) // 手动确认消息
	}
	done <- true
}
