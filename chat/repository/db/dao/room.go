package dao

import (
	"common/repository/db"
	models2 "common/repository/db/models"
	"context"
	"gorm.io/gorm"
)

type RoomDao struct {
	*gorm.DB
}

func NewRoomDao(ctx context.Context) *RoomDao {
	return &RoomDao{db.NewDBClient(ctx)}
}

func NewRoomDaoByDB(db *gorm.DB) *RoomDao {
	return &RoomDao{db}
}

// GetById 根据 id 获取
func (dao *RoomDao) GetById(id uint) (room *models2.Room, err error) {
	err = dao.DB.Model(&models2.RoomMember{}).Where("id=?", id).
		First(&room).Error
	return
}

// UpdateById 根据 id 更新
func (dao *RoomDao) UpdateById(id uint, room *models2.Room) (err error) {
	return dao.DB.Model(&models2.Room{}).Where("id=?", id).
		Updates(&room).Error
}

// Create 创建訊息
func (dao *RoomDao) Create(room *models2.Room) error {
	return dao.DB.Model(&models2.Room{}).Create(&room).Error
}
