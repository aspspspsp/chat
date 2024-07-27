package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func DiscoverServiceWithConsul() (string, error) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return "", err
	}

	services, _, err := client.Health().Service("greeter", "", true, nil)
	if err != nil {
		return "", err
	}

	if len(services) == 0 {
		log.Println("test")
		return "", fmt.Errorf("no healthy instances of service found")
	}

	service := services[0]
	address := service.Service.Address
	port := service.Service.Port

	log.Println(fmt.Sprintf("address:%s, port:%d", address, port))

	return fmt.Sprintf("%s:%d", address, port), nil
}

type RegisterServiceFunc func(*grpc.Server)

func StartGrpc(registerServices []RegisterServiceFunc) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	for _, registerService := range registerServices {
		registerService(s)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
