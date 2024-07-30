package main

import (
	"chat/repository/mq/message_broadcast"
	"chat/repository/mq/message_store"
	"chat/routes"
	"common/repository/db"
	"common/utils"
	"context"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// db初始化
	db.InitMySQL()

	// MQ初始化
	message_store.InitMq(ctx)
	message_broadcast.InitMq()

	routes.SetupWsRoutes()

	r := routes.NewRouter()

	// 启动 HTTP 服务
	err := r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
