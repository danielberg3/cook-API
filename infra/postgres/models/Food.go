package models

import "Cook_API/core/domain"

type Food struct {
	ID   int    `db:"food_id"`
	Name string `db:"food_name"`
	Time int    `db:"food_time"`
}

func (dto Food) ToDomain() *domain.Food {
	return domain.NewFoodWithId(dto.ID, dto.Name, dto.Time)
}
