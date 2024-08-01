package server

import (
	"common/pb"
	"common/repository/db/models"
	"context"
	"member/repository/db/dao"
)

type MemberServer struct {
	pb.UnimplementedMemberServiceServer
}

func (s *MemberServer) GetMember(ctx context.Context, in *pb.GetMemberRequest) (*pb.GetMemberResponse, error) {
	id := uint(in.GetId())
	member, _ := dao.NewMemberDao(ctx).GetById(id)
	pMember := models.ConvertToProto(member)

	return &pb.GetMemberResponse{Member: pMember}, nil
}
