package dao

import (
	"common/repository/db"
	"common/repository/db/models"
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
func (dao *RoomDao) GetById(id uint) (room *models.Room, err error) {
	err = dao.DB.Model(&models.RoomMember{}).Where("id=?", id).
		First(&room).Error
	return
}

// UpdateById 根据 id 更新
func (dao *RoomDao) UpdateById(id uint, room *models.Room) (err error) {
	return dao.DB.Model(&models.Room{}).Where("id=?", id).
		Updates(&room).Error
}

// Create 创建訊息
func (dao *RoomDao) Create(room *models.Room) error {
	return dao.DB.Model(&models.Room{}).Create(&room).Error
}

func (dao *RoomDao) DeleteById(id uint) error {
	return dao.DB.Where("id = ?", id).Delete(&models.RoomMember{}).Error
}
