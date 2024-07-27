package main

import (
	"common/utils"
	"github.com/gin-gonic/gin"
	"member/config"
	"member/routes"
)

func main() {
	// 初始化數據庫
	config.DbInit()

	// Gin 路由设置
	r := gin.Default()
	routes.SetupRoutes(r)

	// 初始化consul
	config.ConsulInit()

	// 启动 HTTP 服务
	err := r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
