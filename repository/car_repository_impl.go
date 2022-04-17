package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type carRepositoryImpl struct {
	DB *sql.DB
}

func newCarRepositoryImpl(db *sql.DB) carRepositoryImpl {
	return carRepositoryImpl{DB: db}
}

func (repo carRepositoryImpl) Insert(ctx context.Context, car entitiy.Cars) (entitiy.Cars, error) {
	script := "INSERT INTO cars(name, merk) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, car.Id, car.Name, car.Merk)
	if err != nil {
		return car, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return car, err
	}
	car.Id = int32(id)
	return car, nil
}

func (repo carRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Cars, error) {
	script := "SELECT id, name, merk FROM  WHERE id = ? LIMIT 3"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	car := entitiy.Cars{}

	if err != nil {
		return car, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&car.Id, &car.Name, &car.Merk)
		return car, nil
	} else {
		//tidak ada
		return car, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo carRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Cars, error) {
	script := "SELECT id, name, merk FROM players"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []entitiy.Cars
	for rows.Next() {
		car := entitiy.Cars{}
		rows.Scan(&car.Id, &car.Name, &car.Merk)
		cars = append(cars, car)
	}
	return cars, nil
}

func (repo carRepositoryImpl) Update(ctx context.Context, car *entitiy.Cars) (*entitiy.Cars, error) {
	script := "SELECT cars name = ?, WHERE id = ?"
	rows, err := repo.DB.PrepareContext(ctx, script)
	if err != nil {
		return car, err
	}
	_, err = rows.ExecContext(ctx, car.Id, car.Name, car.Merk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return car, nil
}

func (repo carRepositoryImpl) Delete(ctx context.Context, id int32) (bool, error) {
	script := "DELETE cars WHERE id = ? LIMIT 3"
	rows, err := repo.DB.PrepareContext(ctx, script)
	if err != nil {
		return false, err
	}
	_, err = rows.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
