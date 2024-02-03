package dicontainer

import "Cook_API/APP/API/endpoints/handlers"

func GetFoodHandler() *handlers.FoodHandlers {
	return handlers.NewFoodHandlers(GetFoodServices())
}
