package primary

import "Cook_API/core/domain"

type FoodManager interface {
	CreateFood(food domain.Food) (int, error)
	FetchFood() ([]domain.Food, error)
	GetFood(foodId int) (domain.Food, error)
	UpdateFood(foodId int, food domain.Food) error
	DeleteFood(id int) error
}
