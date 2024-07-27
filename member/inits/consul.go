package inits

import (
	"common/utils"
)

const (
	serviceName = "greeter"
)

func ConsulInit() {
	// 註冊服務
	go utils.RegisterService("greeter", "127.0.0.1", 50051)
}
