package service

import (
	"chat/repository/db/dao"
	"chat/service/grpc"
	"chat/types"
	"common/repository/db/models"
	"context"
	"log"
	"strconv"
	"sync"
)

var (
	RoomSrvIns  *RoomSrv
	RoomSrvOnce sync.Once
)

type RoomSrv struct {
}

func GetRoomSrv() *RoomSrv {
	RoomSrvOnce.Do(func() {
		RoomSrvIns = &RoomSrv{}
	})
	return RoomSrvIns
}

func (s *RoomSrv) Create(ctx context.Context, req *types.CreateRoomReq) {
	name := req.Name

	room := models.Room{
		Name: name,
	}

	_ = dao.NewRoomDao(ctx).Create(&room)
}

func (s *RoomSrv) Delete(ctx context.Context, req *types.DeleteRoomReq) {
	id := req.Id
	err := dao.NewRoomDao(ctx).DeleteById(id)
	if err != nil {
		return
	}
}

func (s *RoomSrv) AddMember(ctx context.Context, req *types.AddToRoomReq) {
	roomId := req.RoomId
	memberId := req.MemberId

	room, _ := dao.NewRoomMemberDao(ctx).GetById(roomId)
	if room == nil {
		log.Println("房間找不到 id :" + strconv.Itoa(int(roomId)))
		return
	}

	member, _ := grpc.GetMember(memberId)
	if member == nil {
		log.Println("會員找不到 id :" + strconv.Itoa(int(memberId)))
		return
	}

	roomMember := models.RoomMember{
		RoomID:   roomId,
		MemberID: memberId,
	}
	err := dao.NewRoomMemberDao(ctx).Create(&roomMember)
	if err != nil {
		return
	}
}

func (s *RoomSrv) RemoveMember(ctx context.Context, req *types.RemoveToRoomReq) {
	roomId := req.RoomId
	memberId := req.MemberId

	room, _ := dao.NewRoomMemberDao(ctx).GetById(roomId)
	if room == nil {
		log.Println("房間找不到 id :" + strconv.Itoa(int(roomId)))
		return
	}

	member, _ := grpc.GetMember(memberId)
	if member == nil {
		log.Println("會員找不到 id :" + strconv.Itoa(int(memberId)))
		return
	}

	err := dao.NewRoomMemberDao(ctx).DeleteByRoomIdMemberId(roomId, memberId)
	if err != nil {
		return
	}
}

func (s *RoomSrv) ListRoomMembers(roomId uint) {
	//room := roomDao.Get(roomId)
	//if room == nil {
	//	return
	//}

}

func (s *RoomSrv) CreateRoom(ctx context.Context, roomName string, ownerId uint) {
	room := models.Room{
		OwnerId: ownerId,
		Name:    roomName,
	}

	err := dao.NewRoomDao(ctx).Create(&room)
	if err != nil {
		return
	}
}

func (s *RoomSrv) RemoveRoom(ctx context.Context, roomId uint) {
	roomDao := dao.NewRoomDao(ctx)
	room, _ := roomDao.GetById(roomId)
	if room == nil {
		return
	}

	err := roomDao.DeleteById(roomId)
	if err != nil {
		return
	}
}
