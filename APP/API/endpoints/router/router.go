package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() *echo.Echo {
	router := echo.New()
	loadMiddleware(router)

	api := router.Group("/api")
	loadFoodRoutes(api)

	return router

}

func loadMiddleware(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
}
