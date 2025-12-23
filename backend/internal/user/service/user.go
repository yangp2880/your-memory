package service

import (
	"backend/internal/user/repository"
)

type UserService interface{}
type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{repo: repo}
}
