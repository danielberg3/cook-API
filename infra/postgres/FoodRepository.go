package postgres

import (
	"Cook_API/core/domain"
	"Cook_API/core/interfaces/repository"
	"Cook_API/infra/postgres/models"
	"errors"
	"log"
)

var _ repository.FoodLoader = (*FoodRepository)(nil)

type FoodRepository struct {
	Connector
}

func (repo FoodRepository) GetFoodsDB() ([]domain.Food, error) {
	conn, err := repo.getConnection()

	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer repo.closeConnection(conn)

	query := `
		SELECT id as food_id, 
		    name as food_name, 
		    time as food_time 
		from food order by id;
	`

	var foodList []models.Food

	if err = conn.Select(&foodList, query); err != nil {
		log.Print(err)
		return nil, err
	}

	var Foods []domain.Food
	for _, FoodDTO := range foodList {
		Foods = append(Foods, *FoodDTO.ToDomain())
	}

	return Foods, nil
}

func (repo FoodRepository) InsertFoodDB(food domain.Food) (int, error) {
	conn, err := repo.getConnection()

	if err != nil {
		log.Print(err)
		return -1, err
	}
	defer repo.closeConnection(conn)

	query := `INSERT INTO food(name, time) values($1, $2) returning id;`

	var FoodId int
	err = conn.Get(
		&FoodId,
		query,
		food.Name(),
		food.Time(),
	)

	if err != nil {
		log.Print(err)
		return -1, err
	}

	return FoodId, nil
}

func (repo FoodRepository) GetFoodDB(foodId int) (domain.Food, error) {
	conn, err := repo.getConnection()
	if err != nil {
		return domain.Food{}, err
	}
	defer repo.closeConnection(conn)

	row, err := conn.Query("select * from food where id = $1", foodId)
	if err != nil {
		return domain.Food{}, err
	}

	var foodDB models.Food

	if err != nil {
		return domain.Food{}, err
	}

	if row.Next() {
		if err = row.Scan(
			&foodDB.ID,
			&foodDB.Name,
			&foodDB.Time,
		); err != nil {
			return domain.Food{}, err
		}
	} else {
		return domain.Food{}, errors.New("The food not Exist.")
	}
	foodDomain := foodDB.ToDomain()
	return *foodDomain, nil
}

func (repo FoodRepository) DeleteFoodDB(foodId int) error {
	conn, err := repo.getConnection()
	if err != nil {
		return err
	}
	defer repo.closeConnection(conn)

	query := `delete from food where id = $1`
	stmt, err := conn.Prepare(query)
	if err != nil {
		log.Print("Entrou aqui:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(foodId)
	if err != nil {
		return err
	}

	return nil

}

func (repo FoodRepository) UpdateFoodDB(foodId int, food domain.Food) error {
	conn, err := repo.getConnection()
	if err != nil {
		return err
	}
	defer repo.closeConnection(conn)

	query := `update food set name = $1, time = $2 where id = $3;`

	res, err := conn.Exec(query, food.Name(), food.Time(), foodId)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Print(count)

	return nil
}

func NewFoodRepository(connectionManager Connector) *FoodRepository {
	return &FoodRepository{connectionManager}
}
