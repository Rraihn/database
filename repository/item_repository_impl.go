package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type itemRepositoryImpl struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepositoryImpl{DB: db}
}

func (repo *itemRepositoryImpl) Insert(ctx context.Context, item entitiy.Items) (entitiy.Items, error) {
	script := "INSERT INTO items(name, qty) VALUES (?, ?)"
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
	script := "SELECT id, name, qty FROM items WHERE id = ? LIMIT 3"
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
	script := "SELECT id, name, qty FROM items"
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

func (repo *itemRepositoryImpl) Update(ctx context.Context, item entitiy.Items) (entitiy.Items, error) {
	script := "UPDATE items SET name = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, script, item.Name, item.Id)
	if err != nil {
		return item, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return item, err
	}
	if rowCnt == 0 {
		return item, err
	}
	return item, nil
}

func (repo itemRepositoryImpl) Delete(ctx context.Context, item entitiy.Items) (entitiy.Items, error) {
	script := "DELETE FROM items WHERE id = ?"

	result, err := repo.DB.ExecContext(ctx, script, item.Id)
	if err != nil {
		return item, err
	}
	rowCnt, err := result.RowsAffected()
	if rowCnt == 0 {
		return item, err
	}
	return item, err
}
