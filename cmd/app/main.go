package main

import (
	"context"
	"github.com/YANcomp/yanco-backend/internal/app"
)

const configsDir = "configs"

func main() {
	ctx := context.Background()

	app.Run(ctx, configsDir)
}
