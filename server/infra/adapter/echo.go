package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Adapter struct {
	*echo.Echo
}

func NewServer() *Adapter {
	s := echo.New()
	s.Use(middleware.Logger())

	return &Adapter{s}
}
