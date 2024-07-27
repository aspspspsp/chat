package config

import (
	"common/utils"
	"member/models"
)

func DbInit() {
	utils.DbInit("root", "oh_my_ody!", "127.0.0.1", 13306, "chat")
	// 自動遷移
	utils.DB.AutoMigrate(&models.Member{})

	defer utils.Close() // 確保在程序結束時關閉數據庫連接
}
