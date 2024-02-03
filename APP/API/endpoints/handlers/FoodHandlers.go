package handlers

import (
	"Cook_API/APP/API/endpoints/dto/request"
	"Cook_API/APP/API/endpoints/dto/response"
	"Cook_API/core/interfaces/primary"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type FoodHandlers struct {
	FoodService primary.FoodManager
}

func (handler FoodHandlers) PostFood(c echo.Context) error {
	var foodDTO request.Food

	if err := c.Bind(&foodDTO); err != nil {
		return c.JSON(http.StatusBadRequest,
			response.NewError(http.StatusBadRequest,
				"Algo está incorreto em sua requisição",
				err.Error(),
			))
	}

	food := foodDTO.ToDomain()
	foodID, err := handler.FoodService.CreateFood(*food)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Parece que houve um problema em sua base de dados",
				err.Error(),
			))
	}

	return c.JSON(http.StatusCreated, response.Created{ID: foodID})
}

func (handler FoodHandlers) GetFoods(c echo.Context) error {
	foods, err := handler.FoodService.FetchFood()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Parece que algo deu errado em sua base de dados",
				err.Error(),
			))
	}

	var foodListResponse []response.Food
	for _, food := range foods {
		foodListResponse = append(foodListResponse, *response.NewFood(food))
	}

	return c.JSON(http.StatusOK, foodListResponse)
}

func (handler FoodHandlers) GetFood(c echo.Context) error {
	param := c.Param("foodId")
	foodID, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.NewError(http.StatusBadRequest,
				"Erro no parâmetro passado.",
				err.Error(),
			))
	}

	food, err := handler.FoodService.GetFood(foodID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Houve um erro em sua base de dados",
				err.Error(),
			))
	}

	foodDTO := response.NewFood(food)

	return c.JSON(http.StatusOK, foodDTO)
}

func (handler FoodHandlers) DeleteFood(c echo.Context) error {
	param := c.Param("foodId")
	foodId, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.NewError(http.StatusBadRequest,
				"Houve um erro na leitura dos parâmetros da requisição.",
				err.Error(),
			))
	}

	_, err = handler.FoodService.GetFood(foodId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Houve um erro em sua base de dados",
				err.Error(),
			))
	}

	err = handler.FoodService.DeleteFood(foodId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Houve um erro em sua base de dados.",
				err.Error(),
			))
	}

	return c.JSON(http.StatusOK, "Food is deleted!")

}

func (handler FoodHandlers) UpdateFood(c echo.Context) error {
	param := c.Param("foodId")
	foodId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.NewError(http.StatusBadRequest,
				"Houve um erro no parâmetro passado.",
				err.Error(),
			))
	}

	_, err = handler.FoodService.GetFood(foodId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Houve um erro em sua base de dados",
				err.Error(),
			))
	}

	var foodDTO request.Food
	if err = c.Bind(&foodDTO); err != nil {
		return c.JSON(http.StatusBadRequest,
			response.NewError(http.StatusBadRequest,
				"Houve um erro com o corpo da requisição.",
				err.Error(),
			))
	}

	foodDomain := foodDTO.ToDomain()

	if err = handler.FoodService.UpdateFood(foodId, *foodDomain); err != nil {
		return c.JSON(http.StatusInternalServerError,
			response.NewError(http.StatusInternalServerError,
				"Houve um erro em sua base de dados.",
				err.Error(),
			))
	}

	return c.JSON(http.StatusOK, "Comida atualizada com sucesso.")

}

func NewFoodHandlers(FoodService primary.FoodManager) *FoodHandlers {
	return &FoodHandlers{FoodService: FoodService}
}
