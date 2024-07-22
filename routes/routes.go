package routes

import (
	"github.com/labstack/echo"
	"github.com/viay1668/spectre/routes/api"
	"github.com/viay1668/spectre/routes/web"
)

func RegisterRoutes(e *echo.Echo) {

	// grouping routes
	apiRoutes := e.Group("/api")
	appRoutes := e.Group("")
	// Frontend/web routes
	api.RegisterAuthRoutes(apiRoutes)
	web.RegisterAuthRoutes(appRoutes)
	web.RegisterDashboardRoutes(appRoutes)
	// web.RegisterProfileRoutes(web)

	// ... other frontend route registrations

	// API routes
	api.RegisterAuthRoutes(apiRoutes)

}
