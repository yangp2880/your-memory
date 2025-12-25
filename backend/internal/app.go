package internal

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/http"
	"backend/internal/memory"
	"backend/internal/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
	memory.Module,
	config.Module,
	db.Module,
	http.Module,
)
