package types

type CreateRoomReq struct {
	Name string `form:"name"  json:"name" :"name"`
}

type DeleteRoomReq struct {
	Id uint `form:"id"  json:"id" :"id"`
}

type AddToRoomReq struct {
	RoomId   uint `form:"roomId"  json:"roomId" :"roomId"`
	MemberId uint `form:"memberId"  json:"memberId" :"memberId"`
}

type RemoveToRoomReq struct {
	RoomId   uint `form:"roomId"  json:"roomId" :"roomId"`
	MemberId uint `form:"memberId"  json:"memberId" :"memberId"`
}
