package dao

import (
	"gorm.io/gorm"
)

type UserDao interface{}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}
