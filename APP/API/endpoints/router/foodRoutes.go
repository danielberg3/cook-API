package router

import (
	"Cook_API/APP/API/endpoints/dicontainer"
	"github.com/labstack/echo/v4"
)

func loadFoodRoutes(api *echo.Group) {
	foodGroup := api.Group("/food")

	foodHandlers := dicontainer.GetFoodHandler()

	foodGroup.POST("/new", foodHandlers.PostFood)
	foodGroup.GET("", foodHandlers.GetFoods)
	foodGroup.GET("/:foodId", foodHandlers.GetFood)
	foodGroup.PUT("/:foodId", foodHandlers.UpdateFood)
	foodGroup.DELETE("/:foodId", foodHandlers.DeleteFood)

}
