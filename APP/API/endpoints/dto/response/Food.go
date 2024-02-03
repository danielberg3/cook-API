package response

import "Cook_API/core/domain"

type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

func NewFood(food domain.Food) *Food {
	return &Food{
		ID:   food.Id(),
		Name: food.Name(),
		Time: food.Time(),
	}
}
