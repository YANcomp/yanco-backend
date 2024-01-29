package main

import (
	"github.com/YANcomp/yanco-backend/internal/app"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
