package models

import "time"

type Message struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	ChatRoomID uint      `json:"chat_room_id" gorm:"not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at"`
}
