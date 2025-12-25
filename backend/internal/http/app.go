package http

import (
	"backend/internal/http/rounter"
	"backend/internal/http/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(rounter.NewGinEngine),
	fx.Invoke(service.RunHTTPServer),
)
