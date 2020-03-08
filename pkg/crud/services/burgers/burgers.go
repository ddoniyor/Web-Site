package burgers

import (
	"context"
	"crud/pkg/crud/errors"
	"crud/pkg/crud/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BurgersSvc struct {
	pool *pgxpool.Pool
}

func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		panic(errors.Erroring("pool can't be nil"))
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	list = make([]models.Burger, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT id, name, price FROM burgers WHERE removed = FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Burger{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *BurgersSvc) Save(model models.Burger) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors.ApiError("can't execute pool: ", err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO burgers(name, price) VALUES ($1, $2);", model.Name, model.Price)
	if err != nil {
		return errors.ApiError("can't save burger: ", err)
	}
	return nil
}

func (service *BurgersSvc) RemoveById(id int) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors.ApiError("can't execute pool: ", err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "UPDATE burgers SET removed = true where id = $1;",id)
	if err != nil {
		return errors.ApiError("can't remove burger: ", err)
	}
	return nil
}
