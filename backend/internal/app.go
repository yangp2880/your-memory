package internal

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/user"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.ProviderSet,
	fx.Provide(db.InitMysql, config.MysqlConfig),
)
