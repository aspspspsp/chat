package types

type AddToRoomReq struct {
	RoomId   uint `form:"roomId"  json:"roomId" :"content"`
	MemberId uint `form:"memberId"  json:"memberId" :"content"`
}

type RemoveToRoomReq struct {
	RoomId   uint `form:"roomId"  json:"roomId" :"content"`
	MemberId uint `form:"memberId"  json:"memberId" :"content"`
}
