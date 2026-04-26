package dao

import (
	"ZpcGoProject/config"
	"ZpcGoProject/internal/model"
)

type UserDao struct{}

func (u *UserDao) Create(user *model.User) error {
	return config.DB.Create(user).Error
}
func (u *UserDao) GetByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return user, err
}
func (u *UserDao) List() ([]model.User, error) {
	var list []model.User
	err := config.DB.Find(&list).Error
	return list, err
}
func (u *UserDao) Update(user *model.User) error {
	return config.DB.Save(user).Error
}
func (u *UserDao) Delete(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
