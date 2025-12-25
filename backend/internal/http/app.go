package http

import "go.uber.org/fx"

var module = fx.Options(
	fx.Provide(),
	fx.Invoke(),
)
