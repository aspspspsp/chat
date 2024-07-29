package models

type RoomMember struct {
	ID       uint `json:"id" gorm:"primary_key"`
	RoomID   uint `json:"room_id" gorm:"not null"`
	MemberID uint `json:"member_id" gorm:"not null"`
}
