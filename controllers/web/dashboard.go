package web

import (
	"context"

	"github.com/labstack/echo"
	"github.com/viay1668/spectre/templates"
)

func RenderDashboard(c echo.Context) error {
	dashboardComponent := templates.Index()
	return dashboardComponent.Render(context.Background(), c.Response())
}
