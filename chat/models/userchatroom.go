package models

type UserChatRoom struct {
	ID         uint `json:"id" gorm:"primary_key"`
	UserID     uint `json:"user_id" gorm:"not null"`
	ChatRoomID uint `json:"chat_room_id" gorm:"not null"`
}
