package dao

import (
	"chat/repository/db/models"
	"common/repository/db"
	"context"
	"gorm.io/gorm"
)

type MessageDao struct {
	*gorm.DB
}

func NewMessageDao(ctx context.Context) *MessageDao {
	return &MessageDao{db.NewDBClient(ctx)}
}

// GetById 根据 id 获取訊息
func (dao *MessageDao) GetById(id uint) (message *models.Message, err error) {
	err = dao.DB.Model(&models.Message{}).Where("id=?", id).
		First(&message).Error
	return
}

// UpdateById 根据 id 更新訊息
func (dao *MessageDao) UpdateById(id uint, message *models.Message) (err error) {
	return dao.DB.Model(&models.Message{}).Where("id=?", id).
		Updates(&message).Error
}

// Create 创建訊息
func (dao *MessageDao) Create(message *models.Message) error {
	return dao.DB.Model(&models.Message{}).Create(&message).Error
}
