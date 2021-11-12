package server

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/takeout-users-api/server/infra/adapter"
	"net/http"
)

type HealthCheckRouter struct {
	fw *adapter.Adapter
}

func NewHealthCheck(fw *adapter.Adapter) *HealthCheckRouter {
	hc := &HealthCheckRouter{
		fw: fw,
	}

	hc.HealthCheck()

	return hc
}

func (h *HealthCheckRouter) HealthCheck() {
	h.fw.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})
}
