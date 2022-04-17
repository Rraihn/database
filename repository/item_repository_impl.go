package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type itemRepositoryImpl struct {
	DB sql.DB
}

func (repo *itemRepositoryImpl) Insert(ctx context.Context, item entitiy.Items) (entitiy.Items, error) {
	script := "INSERT INTO comments(name, qty) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, item.Name, item.Qty)
	if err != nil {
		return item, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}
	item.Id = int32(id)
	return item, nil
}

func (repo *itemRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Items, error) {
	script := "SELECT id, name, qty FROM  WHERE id = ? LIMIT 3"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	item := entitiy.Items{}

	if err != nil {
		return item, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&item.Id, &item.Name, &item.Qty)
		return item, nil
	} else {
		//tidak ada
		return item, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *itemRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Items, error) {
	script := "SELECT id, name, qty FROM item"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []entitiy.Items
	for rows.Next() {
		item := entitiy.Items{}
		rows.Scan(&item.Id, &item.Name, &item.Qty)
		items = append(items, item)
	}
	return items, nil
}

func (repo *itemRepositoryImpl) Update(ctx context.Context, item *entitiy.Items) (*entitiy.Items, error) {
	script := "SELECT item Name = ?, WHERE id = ?"
	rows, err := repo.DB.PrepareContext(ctx, script)
	if err != nil {
		return item, err
	}
	_, err = rows.ExecContext(ctx, item.Id, item.Id, item.Qty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return item, nil
}
