package config

import (
	"chat/models"
	"common/utils"
)

func DbInit() {
	utils.DbInit("root", "oh_my_ody!", "127.0.0.1", 13306, "chat")

	// 自動遷移
	utils.DB.AutoMigrate(&models.ChatRoom{})
	utils.DB.AutoMigrate(&models.Message{})
	utils.DB.AutoMigrate(&models.UserChatRoom{})

	defer utils.Close() // 確保在程序結束時關閉數據庫連接
}
