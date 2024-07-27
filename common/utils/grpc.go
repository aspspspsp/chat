package utils

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type RegisterServiceFunc func(*grpc.Server)
type CallGRPCServiceFunc func(ctx context.Context, clientConn *grpc.ClientConn) (interface{}, error)

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
		return "", fmt.Errorf("no healthy instances of service found")
	}

	service := services[0]
	address := service.Service.Address
	port := service.Service.Port

	log.Println(fmt.Sprintf("address:%s, port:%d", address, port))

	return fmt.Sprintf("%s:%d", address, port), nil
}

func StartGrpc(port int, registerServices []RegisterServiceFunc) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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

func CallGRPCService(serviceAddress string, callFunc CallGRPCServiceFunc, retries int, delay time.Duration) (interface{}, error) {
	var lastErr error
	for i := 0; i < retries; i++ {
		conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			lastErr = err
			time.Sleep(delay)
			continue
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Printf("Failed to close connection: %v", err)
			}
		}(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		result, err := callFunc(ctx, conn)
		if err == nil {
			return result, nil
		}

		lastErr = err
		time.Sleep(delay)
	}

	return nil, lastErr
}
