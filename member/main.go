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
	name := "world"
	serviceAddress, err := utils.DiscoverServiceWithConsul()
	if err != nil {
		log.Fatalf("did not connect1: %v", err)
	}

	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect2: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("fffff: %v", err)
		}
	}(conn)

	c := pb.NewGreeterClient(conn)
	name = "world"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	hello, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", hello.GetMessage())

	// 启动 HTTP 服务
	err = r.Run(":8080")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
