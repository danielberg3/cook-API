package request

import "Cook_API/core/domain"

type Food struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Time int    `json:"Time"`
}

func (dto Food) ToDomain() *domain.Food {
	return domain.NewFood(dto.Name, dto.Time)
}
