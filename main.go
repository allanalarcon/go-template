package main

import (
	"bitbucket.org/devsu/go-template/app"

	_ "bitbucket.org/devsu/go-template/docs"
)

// @title Swagger Example API
// @version 1.0
// @description GoTemplate API
// @title GoTemplate API

// @host localhost:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := app.New()
	app.LoadConfig()
	app.Initialize()
	app.Run("0.0.0.0:8585")
}
