package main

import (
	"common/pb"
	"common/repository/cache"
	"common/repository/db"
	"common/repository/rpc"
	"common/utils"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"member/handlers"
	"member/inits"
	"member/routes"
	"time"
)

func main() {
	cache.InitRedis()

	db.InitMySQL()
	// 自動遷移
	//db.DB.AutoMigrate(&models.Member{})
	//defer db.Close() // 確保在程序結束時關閉數據庫連接

	// Gin 路由设置
	r := gin.Default()
	routes.SetupRoutes(r)

	// 註冊至服務發現
	inits.ConsulInit()

	// GRPC 註冊
	registerServices := []rpc.RegisterServiceFunc{
		func(s *grpc.Server) {
			pb.RegisterGreeterServer(s, &handlers.Server{})
		},
	}
	inits.GrpcInit(registerServices)

	// grpc 調用
	serviceAddress, err := rpc.DiscoverServiceWithConsul()
	if err != nil {
		log.Fatalf("did not connect1: %v", err)
	}
	maxRetries := 3
	retryInterval := 2 * time.Second
	result, err := rpc.CallGRPCService(serviceAddress, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewGreeterClient(conn)
		return client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	}, maxRetries, retryInterval)

	helloReply, ok := result.(*pb.HelloReply)
	if !ok {
		log.Fatalf("Unexpected response types: %T", result)
	}
	log.Println(helloReply)

	// 启动 HTTP 服务
	err = r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
