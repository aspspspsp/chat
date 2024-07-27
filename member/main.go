package main

import (
	"common/utils"
	"github.com/gin-gonic/gin"
	"member/database"
	"member/routes"
)

func main() {
	// 初始化數據庫
	database.Init()
	defer database.Close() // 確保在程序結束時關閉數據庫連接

	// Gin 路由设置
	r := gin.Default()

	routes.SetupRoutes(r)

	go utils.RegisterService("member", "localhost", 50051)

	// 启动 HTTP 服务
	err := r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
