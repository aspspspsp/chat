package inits

import (
	"common/utils"
)

func GrpcInit(registerServices []utils.RegisterServiceFunc) {
	go utils.StartGrpc(50051, registerServices)
}
