package main

import (
	"common/pb"
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
	// 初始化數據庫
	inits.DbInit()

	// Gin 路由设置
	r := gin.Default()
	routes.SetupRoutes(r)

	// 註冊至服務發現
	inits.ConsulInit()

	// GRPC 註冊
	registerServices := []utils.RegisterServiceFunc{
		func(s *grpc.Server) {
			pb.RegisterGreeterServer(s, &handlers.Server{})
		},
	}
	inits.GrpcInit(registerServices)

	// grpc 調用
	serviceAddress, err := utils.DiscoverServiceWithConsul()
	if err != nil {
		log.Fatalf("did not connect1: %v", err)
	}
	maxRetries := 3
	retryInterval := 2 * time.Second
	result, err := utils.CallGRPCService(serviceAddress, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewGreeterClient(conn)
		return client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	}, maxRetries, retryInterval)

	helloReply, ok := result.(*pb.HelloReply)
	if !ok {
		log.Fatalf("Unexpected response type: %T", result)
	}
	log.Println(helloReply)

	// 启动 HTTP 服务
	err = r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
