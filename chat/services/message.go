package services

import (
	"chat/repository/db/models"
	"chat/repository/mq/message_broadcast"
	"chat/types"
	"sync"
)

var (
	MessageSrvIns  *MessageSrv
	MessageSrvOnce sync.Once
)

type MessageSrv struct {
}

func GetMessageSrv() *MessageSrv {
	MessageSrvOnce.Do(func() {
		MessageSrvIns = &MessageSrv{}
	})
	return MessageSrvIns
}

// SendMessage 發送訊息
func (s *MessageSrv) SendMessage(req *types.SendMessageReq) (resp interface{}, err error) {
	content := req.Content

	// 初始化 channel 和 message
	message := models.Message{ // 假设 types.Message 是你定义的消息结构体
		Content: content,
	}

	message_broadcast.PublishMessage(message)

	return
}
