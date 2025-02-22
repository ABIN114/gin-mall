package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById 根据Id获取用户
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// UpdateUserById 根据Id更新用户信息

func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error

}
