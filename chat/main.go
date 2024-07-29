package main

import (
	"chat/routes"
	"common/repository/db"
	"common/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// db初始化
	db.InitMySQL()

	// Gin 路由设置
	r := gin.Default()

	routes.SetupWsRoutes()

	// 启动 HTTP 服务
	err := r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
