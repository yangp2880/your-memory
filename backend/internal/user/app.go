package user

import (
	"backend/internal/user/controller"
	"backend/internal/user/repository"
	"backend/internal/user/repository/dao"
	"backend/internal/user/service"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	controller.NewUserController,
	service.NewUserService,
	repository.NewUserRepo,
	dao.NewUserDao,
)
