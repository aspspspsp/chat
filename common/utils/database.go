package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var err error
var Node *snowflake.Node

func DbInit(dbUser string, dbPass string, dbUrl string, dbPort int, dbName string) {
	// 初始化雪花算法節點
	Node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatal("Failed to initialize snowflake node:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbUrl, dbPort, dbName)

	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to connect to database")
	}
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return
		}
	}
}
