package service

import (
	"chat/repository/mq/message_broadcast"
	"chat/types"
	"common/repository/db/models"
	"sync"
	"time"
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

	// 初始化 channel 和 message
	message := models.Message{ // 假设 types.Message 是你定义的消息结构体
		RoomId:    req.RoomId,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	message_broadcast.PublishMessage(message)

	return
}
