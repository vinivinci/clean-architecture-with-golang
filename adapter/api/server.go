package api

import (
	"net/http"
	repository "test-architecture/entity"
	"test-architecture/usecase"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Repository repository.UserRepository
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.POST("/user", w.Process)
	e.Logger.Fatal(e.Start(":3000"))
}

func (w WebServer) Process(c echo.Context) error {
	userDto := &usecase.CreateUserDTOInput{}
	c.Bind(userDto)
	usecaseInjection := usecase.NewCreateUser(w.Repository)
	output, _ := usecaseInjection.Execute(*userDto)

	return c.JSON(http.StatusCreated, output)
}
