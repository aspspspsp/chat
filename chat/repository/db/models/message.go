package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	RoomId    uint      `json:"room_id" gorm:"not null"`
	MemberID  uint      `json:"member_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"types:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
