package dicontainer

import (
	"Cook_API/core/interfaces/repository"
	"Cook_API/infra/postgres"
)

func GetFoodRepository() repository.FoodLoader {
	return postgres.NewFoodRepository(GetPSQLConectionManager())
}

func GetPSQLConectionManager() postgres.Connector {
	return &postgres.DatabaseConnectionManager{}
}
