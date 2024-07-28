package inits

import (
	"common/repository/rpc"
)

func GrpcInit(registerServices []rpc.RegisterServiceFunc) {
	go rpc.InitGrpc(50051, registerServices)
}
