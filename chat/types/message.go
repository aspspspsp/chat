package types

type SendMessageReq struct {
	RoomId  uint   `form:"roomId"  json:"roomId" :"content"`
	Content string `form:"content"  json:"content" :"content"`
}
