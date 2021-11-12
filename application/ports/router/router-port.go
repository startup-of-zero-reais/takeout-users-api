package router

import (
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(c echo.Context) error

type Router interface {
	Create(route string)
	Update(route string)
	GetOne(route string)
	List(route string)
	Delete(route string)
}
