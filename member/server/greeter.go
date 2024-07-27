package server

import (
	"common/pb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) sayHello(in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
