package main

import (
	"common/pb"
	"common/repository/cache"
	"common/repository/db"
	"common/repository/rpc"
	"common/utils"
	"google.golang.org/grpc"
	"member/inits"
	"member/routes"
	"member/server"
)

func main() {
	cache.InitRedis()

	db.InitMySQL()
	// 自動遷移
	//db.DB.AutoMigrate(&models.Member{})
	//defer db.Close() // 確保在程序結束時關閉數據庫連接

	// Gin 路由设置
	r := routes.NewRouter()

	// 註冊至服務發現
	inits.ConsulInit()

	// GRPC 註冊
	registerServices := []rpc.RegisterServiceFunc{
		func(s *grpc.Server) {
			pb.RegisterGreeterServer(s, &server.Server{})
		},
		// 會員GRPC server
		func(s *grpc.Server) {
			pb.RegisterMemberServiceServer(s, &server.MemberServer{})
		},
	}
	inits.GrpcInit(registerServices)

	// 启动 HTTP 服务
	err := r.Run(":8081")
	if err != nil {
		utils.FailOnError(err, "member服務啟動失敗")
		return
	}
}
