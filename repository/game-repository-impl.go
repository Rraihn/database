package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type gameRepositoryImpl struct {
	DB *sql.DB
}

func NewGameRepositoryImpl(db *sql.DB) gameRepositoryImpl {
	return gameRepositoryImpl{DB: db}
}

func (repo gameRepositoryImpl) Insert(ctx context.Context, game entitiy.Games) (entitiy.Games, error) {
	script := "INSERT INTO games(name, genre) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, game.Name, game.Genre)
	if err != nil {
		return game, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return game, err
	}
	game.Id = int32(id)
	return game, nil
}

func (repo gameRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Games, error) {
	script := "SELECT id, name, genre FROM games WHERE id = ? LIMIT 3"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	game := entitiy.Games{}

	if err != nil {
		return game, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&game.Id, &game.Name, &game.Genre)
		return game, nil
	} else {
		//tidak ada
		return game, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo gameRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Games, error) {
	script := "SELECT id, name, genre FROM games"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var games []entitiy.Games
	for rows.Next() {
		game := entitiy.Games{}
		rows.Scan(&game.Id, &game.Name, &game.Genre)
		games = append(games, game)
	}
	return games, nil
}

func (repo gameRepositoryImpl) Update(ctx context.Context, game entitiy.Games) (entitiy.Games, error) {
	script := "UPDATE games SET name = ?, genre = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, script, game.Name, game.Genre, game.Id)
	if err != nil {
		return game, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return game, err
	}
	if rowCnt == 0 {
		return game, err
	}
	return game, nil
}

func (repo gameRepositoryImpl) Delete(ctx context.Context, game entitiy.Games) (entitiy.Games, error) {
	script := "DELETE FROM games WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, script, game.Id)
	if err != nil {
		return game, err
	}
	rowCnt, err := result.RowsAffected()
	if rowCnt == 0 {
		return game, err
	}
	return game, err
}
