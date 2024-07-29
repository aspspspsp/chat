package models

import (
	"common/configs"
	"common/consts"
	conf "github.com/CocaineCong/gin-mall/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// Member 會員數據結構
type Member struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Nickname  string
	Avatar    string `gorm:"size:1000"`
	Status    string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

// CheckPassword 校验密码
func (u *Member) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// AvatarURL 头像地址
func (u *Member) AvatarURL() string {
	if configs.GetConfig().System.UploadModel == consts.UploadModelOss {
		return u.Avatar
	}
	pConfig := conf.Config.PhotoPath
	return pConfig.PhotoHost + conf.Config.System.HttpPort + pConfig.AvatarPath + u.Avatar
}
