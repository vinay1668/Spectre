package web

import (
	"github.com/labstack/echo"
	dashboard_controller "github.com/viay1668/spectre/controllers/web"
)

func RegisterDashboardRoutes(e *echo.Group) {
	e.GET("/", dashboard_controller.RenderDashboard)
}
