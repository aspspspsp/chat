package loadbalancer

import (
	"common/configs"
	"github.com/hashicorp/consul/api"
	"log"
)

/**
服務註冊
*/

func InitConsul() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	cConfig := configs.GetConfig().Consul

	registration := &api.AgentServiceRegistration{
		ID:      cConfig.ServiceName,
		Name:    cConfig.ServiceName,
		Tags:    []string{"q1mi", "hello"}, // 为服务打标签
		Address: cConfig.ServiceAddr,
		Port:    cConfig.ServicePort,
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
