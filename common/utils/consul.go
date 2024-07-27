package utils

import (
	"github.com/hashicorp/consul/api"
	"log"
)

/**
服務註冊
*/

func RegisterService(serviceName, serviceAddress string, servicePort int) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceName,
		Name:    serviceName,
		Tags:    []string{"q1mi", "hello"}, // 为服务打标签
		Address: serviceAddress,
		Port:    servicePort,
		// TODO: 健康檢查策略
		//Check: &api.AgentServiceCheck{
		//	GRPC:                           serviceAddress + ":" + string(rune(servicePort)),
		//	Interval:                       "10s",
		//	Timeout:                        "1s",
		//	DeregisterCriticalServiceAfter: "1m", // 如果检查失败超过1分钟，注销服务
		//},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}
}

/*
服務發現
*/
