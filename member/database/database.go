package database

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"member/models"
)

var DB *gorm.DB
var err error
var Node *snowflake.Node

func Init() {
	// 初始化雪花算法節點
	Node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatal("Failed to initialize snowflake node:", err)
	}

	DB, err = gorm.Open("mysql", "root:oh_my_ody!@tcp(127.0.0.1:13306)/chat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to connect to database")
	}

	// 自動遷移
	DB.AutoMigrate(&models.Member{})
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return
		}
	}
}
