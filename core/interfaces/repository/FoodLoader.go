package repository

import "Cook_API/core/domain"

type FoodLoader interface {
	InsertFoodDB(food domain.Food) (int, error)
	GetFoodsDB() ([]domain.Food, error)
	GetFoodDB(foodId int) (domain.Food, error)
	UpdateFoodDB(foodId int, food domain.Food) error
	DeleteFoodDB(id int) error
}
