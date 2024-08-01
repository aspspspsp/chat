package grpc

import (
	"common/pb"
	"common/repository/db/models"
	"common/repository/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetMember(id uint) (*models.Member, error) {
	serviceAddress, err := rpc.DiscoverServiceWithConsul()
	if err != nil {
		log.Println("Service discovery failed")
		return nil, err
	}

	maxRetries := 3
	retryInterval := 2 * time.Second

	result, err := rpc.CallGRPCService(serviceAddress, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewMemberServiceClient(conn)
		return client.GetMember(ctx, &pb.GetMemberRequest{Id: int32(id)})
	}, maxRetries, retryInterval)

	getMemberReply, ok := result.(*pb.GetMemberResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	pbMember := getMemberReply.GetMember()
	member, _ := models.ConvertFromProto(pbMember)

	return &member, nil
}
