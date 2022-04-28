package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repo *userRepositoryImpl) Insert(ctx context.Context, user entitiy.Users) (entitiy.Users, error) {
	script := "INSERT INTO users(username, password) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, user.Username, user.Password)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.Id = int32(id)
	return user, nil
}

func (repo *userRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Users, error) {
	script := "SELECT id, username, password FROM users WHERE id = ?"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	user := entitiy.Users{}

	if err != nil {
		return user, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&user.Id, &user.Username, &user.Password)
		return user, nil
	} else {
		//tidak ada
		return user, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *userRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Users, error) {
	script := "SELECT id, username, password FROM users"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entitiy.Users
	for rows.Next() {
		user := entitiy.Users{}
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepositoryImpl) Update(ctx context.Context, user entitiy.Users) (entitiy.Users, error) {
	script := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, script, user.Username, user.Password, user.Id)
	if err != nil {
		return user, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return user, err
	}
	if rowCnt == 0 {
		return user, err
	}
	return user, nil
}

func (repo *userRepositoryImpl) Delete(ctx context.Context, users entitiy.Users) (entitiy.Users, error) {
	script := "DELETE FROM users WHERE id =?"
	result, err := repo.DB.ExecContext(ctx, script, users.Id)
	if err != nil {
		return users, err
	}
	rowCnt, err := result.RowsAffected()
	if rowCnt == 0 {
		return users, err
	}
	return users, nil
}
