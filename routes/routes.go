package routes

import (
	"bitbucket.org/devsu/go-template/controllers"
	"github.com/labstack/echo/v4"
)

type Routes struct{}

func New() *Routes {
	return &Routes{}
}

func (routes *Routes) RegisterV1(e *echo.Group) {
	v1 := e.Group("/v1")

	sample := v1.Group("/sample")
	registerSampleController(sample)

	user := v1.Group("/users")
	registerUserController(user)
}

func registerSampleController(e *echo.Group) {
	sampleController := controllers.NewSampleController()
	e.GET("/hello-world", sampleController.GetHelloWorld)
	e.GET("/users/:id", sampleController.GetUserByID)
}

func registerUserController(e *echo.Group) {

}
