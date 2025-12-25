package memory

import (
	"backend/internal/memory/controller"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controller.NewMemoryController),
)
