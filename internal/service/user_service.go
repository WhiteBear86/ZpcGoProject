package service

import (
	"ZpcGoProject/internal/dao"
	"ZpcGoProject/internal/model"
)

var userDao = dao.UserDao{}

func Create(u *model.User) error          { return userDao.Create(u) }
func GetByID(id uint) (model.User, error) { return userDao.GetByID(id) }
func List() ([]model.User, error)         { return userDao.List() }
func Update(u *model.User) error          { return userDao.Update(u) }
func Delete(id uint) error                { return userDao.Delete(id) }
