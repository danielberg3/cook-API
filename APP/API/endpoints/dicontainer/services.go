package dicontainer

import (
	"Cook_API/core/interfaces/primary"
	"Cook_API/core/services"
)

func GetFoodServices() primary.FoodManager {
	return services.NewFoodServices(GetFoodRepository())
}
