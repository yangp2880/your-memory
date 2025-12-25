package main

import (
	"backend/internal"

	"go.uber.org/fx"
)

func main() {
	fx.New(internal.Module).Run()
}
