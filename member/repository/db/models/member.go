package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Member 會員數據結構
type Member struct {
	gorm.Model
	ID       snowflake.ID `json:"id" gorm:"primary_key"`
	Username string       `json:"username" gorm:"unique"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Email    string       `json:"email" gorm:"unique"`
	Nickname string
	Avatar   string `gorm:"size:1000"`
	Status   string
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 设置密码
func (u *Member) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
