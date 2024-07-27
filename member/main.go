package main

import (
	"github.com/gin-gonic/gin"
	"log"
	registration "member/consul"
	"member/database"
	"member/routes"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 初始化數據庫
	database.Init()
	defer database.Close() // 確保在程序結束時關閉數據庫連接

	// Gin 路由设置
	r := gin.Default()

	routes.SetupRoutes(r)

	go registration.RegisterService("greeter", "localhost", 50051)

	// 启动 HTTP 服务
	r.Run(":8080")
}
