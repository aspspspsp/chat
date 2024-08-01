package dao

import (
	"common/repository/db"
	"common/repository/db/models"
	"context"
	"gorm.io/gorm"
)

type MemberDao struct {
	*gorm.DB
}

func NewMemberDao(ctx context.Context) *MemberDao {
	return &MemberDao{db.NewDBClient(ctx)}
}

func NewMemberDaoByDB(db *gorm.DB) *MemberDao {
	return &MemberDao{db}
}

// GetById 根据 id 获取用户
func (dao *MemberDao) GetById(id uint) (member *models.Member, err error) {
	err = dao.DB.Model(&models.Member{}).Where("id=?", id).
		First(&member).Error
	return
}

// UpdateById 根据 id 更新用户信息
func (dao *MemberDao) UpdateById(id uint, member *models.Member) (err error) {
	return dao.DB.Model(&models.Member{}).Where("id=?", id).
		Updates(&member).Error
}

// Create 创建用户
func (dao *MemberDao) Create(member *models.Member) error {
	return dao.DB.Model(&models.Member{}).Create(&member).Error
}

// IsExistsByUsername 根据username判断是否存在该名字
func (dao *MemberDao) IsExistsByUsername(username string) (member *models.Member, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&models.Member{}).Where("username = ?", username).Count(&count).Error
	if count == 0 {
		return member, false, err
	}
	err = dao.DB.Model(&models.Member{}).Where("username = ?", username).First(&member).Error
	if err != nil {
		return member, false, err
	}
	return member, true, nil
}
