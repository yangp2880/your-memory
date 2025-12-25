package repository

import "backend/internal/user/repository/dao"

type UserRepo interface {
}
type userRepo struct {
	dao dao.UserDao
}

func NewUserRepo(dao dao.UserDao) UserRepo {
	return &userRepo{
		dao: dao,
	}
}
