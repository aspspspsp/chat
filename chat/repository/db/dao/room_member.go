package dao

import (
	"chat/repository/db/models"
	"common/repository/db"
	"context"
	"gorm.io/gorm"
)

type RoomMemberDao struct {
	*gorm.DB
}

func NewRoomMemberDao(ctx context.Context) *RoomMemberDao {
	return &RoomMemberDao{db.NewDBClient(ctx)}
}

// GetById 根据 id 获取
func (dao *RoomMemberDao) GetById(id uint) (roomMember *models.RoomMember, err error) {
	err = dao.DB.Model(&models.RoomMember{}).Where("id=?", id).
		First(&roomMember).Error
	return
}

// UpdateById 根据 id 更新
func (dao *RoomMemberDao) UpdateById(id uint, roomMember *models.RoomMember) (err error) {
	return dao.DB.Model(&models.RoomMember{}).Where("id=?", id).
		Updates(&roomMember).Error
}

// Create 创建訊息
func (dao *RoomMemberDao) Create(roomMember *models.RoomMember) error {
	return dao.DB.Model(&models.RoomMember{}).Create(&roomMember).Error
}
