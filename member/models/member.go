package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
)

// Member 會員數據結構
type Member struct {
	gorm.Model
	ID       snowflake.ID `json:"id" gorm:"primary_key"`
	Username string       `json:"username" gorm:"unique"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Email    string       `json:"email" gorm:"unique"`
}
