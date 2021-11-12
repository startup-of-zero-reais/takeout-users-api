package server

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/takeout-users-api/application/domains/user"
	"github.com/startup-of-zero-reais/takeout-users-api/application/ports/service"
	"github.com/startup-of-zero-reais/takeout-users-api/server/infra/adapter"
	svc "github.com/startup-of-zero-reais/takeout-users-api/server/service/user"
	"net/http"
)

type UserController struct {
	group   *echo.Group
	fw      *adapter.Adapter
	Service service.User
}

func NewUserController(fw *adapter.Adapter) *UserController {
	ur := &UserController{
		group:   fw.Group("users"),
		fw:      fw,
		Service: svc.NewUserService(),
	}

	ur.Create("")
	ur.GetOne("/:id")
	ur.List("")
	ur.Update("/:id")
	ur.Delete("/:id")

	return ur
}

func (ur *UserController) Create(route string) {
	ur.group.POST(route, func(c echo.Context) error {
		u := new(user.User)

		if err := c.Bind(u); err != nil {
			c.Logger().Fatalf(err.Error())
		}

		userCreated, err := ur.Service.Create(u)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, userCreated)
	})
}

func (ur *UserController) Update(route string) {
	ur.group.PUT(route, func(c echo.Context) error {
		//ur.Service.Update()
		return c.JSON(http.StatusOK, map[string]string{})
	})
}

func (ur *UserController) GetOne(route string) {
	ur.group.GET(route, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})
}

func (ur *UserController) List(route string) {
	ur.group.GET(route, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})
}

func (ur *UserController) Delete(route string) {
	ur.group.GET(route, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})
}
