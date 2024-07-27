package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
	"time"
)

func RegisterService(serviceName, serviceAddress string, servicePort int) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceName,
		Name:    serviceName,
		Address: serviceAddress,
		Port:    servicePort,
		Check: &api.AgentServiceCheck{
			GRPC:     serviceAddress + ":" + string(servicePort),
			Interval: "10s",
			Timeout:  "1s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}

	// Keep the service registered
	for {
		time.Sleep(time.Second)
	}
}
