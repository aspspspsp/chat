package inits

import (
	"common/repository/loadbalancer"
)

const (
	serviceName = "greeter"
)

func ConsulInit() {
	// 註冊服務
	go loadbalancer.InitConsul()
}
