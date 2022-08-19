package controllers

import (
	"net/http"
	"strconv"

	"bitbucket.org/devsu/go-template/services"
	"github.com/labstack/echo/v4"
)

type SampleController struct {
	sampleService services.SampleService
}

func NewSampleController() *SampleController {
	sampleService := services.NewSampleServiceImpl()
	controller := &SampleController{
		sampleService: sampleService,
	}
	return controller
}

// GetHelloWorld godoc
// @Summary Get a hello world string
// @Description Get a hello world string
// @ID get-hello-world
// @Tags sample
// @Accept  json
// @Produce  json
// @Router /v1/sample/hello-world [get]
func (c *SampleController) GetHelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, c.sampleService.GetHelloWorld())
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @ID get-user-by-id
// @Tags sample
// @Accept  json
// @Produce  json
// @Param id path number true "User ID"
// @Success 200 {object} dto.UserDto
// @Router /v1/sample/users/{id} [get]
func (c *SampleController) GetUserByID(ctx echo.Context) error {
	print(ctx.ParamNames())
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return err
	}

	user, err := c.sampleService.GetUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}
