package services

import (
	"Cook_API/core/domain"
	"Cook_API/core/interfaces/primary"
	"Cook_API/core/interfaces/repository"
)

var _ primary.FoodManager = (*FoodServices)(nil)

type FoodServices struct {
	FoodRepository repository.FoodLoader
}

func (service FoodServices) CreateFood(food domain.Food) (int, error) {
	foodId, err := service.FoodRepository.InsertFoodDB(food)
	if err != nil {
		return -1, err
	}

	return foodId, nil
}

func (service FoodServices) FetchFood() ([]domain.Food, error) {
	arrayFood, err := service.FoodRepository.GetFoodsDB()

	if err != nil {
		return nil, err
	}

	return arrayFood, nil
}

func (service FoodServices) GetFood(foodId int) (domain.Food, error) {
	food, err := service.FoodRepository.GetFoodDB(foodId)
	if err != nil {
		return domain.Food{}, err
	}

	return food, nil
}

func (service FoodServices) DeleteFood(foodID int) error {
	err := service.FoodRepository.DeleteFoodDB(foodID)
	if err != nil {
		return err
	}

	return nil
}

func (service FoodServices) UpdateFood(foodId int, food domain.Food) error {
	if err := service.FoodRepository.UpdateFoodDB(foodId, food); err != nil {
		return err
	}

	return nil
}

func NewFoodServices(bookRepo repository.FoodLoader) *FoodServices {
	return &FoodServices{bookRepo}
}
