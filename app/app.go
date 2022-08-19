package app

import (
	"log"

	"bitbucket.org/devsu/go-template/router"
	"bitbucket.org/devsu/go-template/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	router *echo.Echo
}

func New() *App {
	r := router.New()
	app := &App{
		router: r,
	}
	return app
}

func (app *App) Initialize() {
	router := app.router

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	apiRouter := router.Group("/api")
	apiRoutes := routes.New()
	apiRoutes.RegisterV1(apiRouter)
}

func (app *App) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error getting env, not comming through %v", err)
		}
	}
}

func (app *App) Run(address string) {
	router := app.router
	router.Logger.Fatal(router.Start(address))
}
