package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type playerRepositoryImpl struct {
	DB *sql.DB
}

func NewPlayerRepository(db *sql.DB) PlayerRepository {
	return &playerRepositoryImpl{DB: db}
}

func (repo *playerRepositoryImpl) Insert(ctx context.Context, player entitiy.Players) (entitiy.Players, error) {
	script := "INSERT INTO players(name, nickname, gender) VALUES (?, ?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, player.Name, player.Nickname, player.Gender)
	if err != nil {
		return player, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return player, err
	}
	player.Id = int32(id)
	return player, nil
}

func (repo *playerRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Players, error) {
	script := "SELECT id, name, nickname, gender FROM players WHERE id = ?"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	player := entitiy.Players{}

	if err != nil {
		return player, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&player.Id, &player.Name, &player.Nickname, &player.Gender)
		return player, nil
	} else {
		//tidak ada
		return player, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *playerRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Players, error) {
	script := "SELECT id, name, nickname, gender FROM players"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []entitiy.Players
	for rows.Next() {
		player := entitiy.Players{}
		rows.Scan(&player.Id, &player.Name, &player.Nickname, &player.Gender)
		players = append(players, player)
	}
	return players, nil
}

func (repo *playerRepositoryImpl) Update(ctx context.Context, player entitiy.Players) (entitiy.Players, error) {
	script := "UPDATE players SET name = ?, nickname = ?, gender = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, script, player.Name, player.Nickname, player.Gender, player.Id)
	if err != nil {
		return player, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return player, nil
	}
	if rowCnt == 0 {
		return player, err
	}
	return player, nil
}

func (repo *playerRepositoryImpl) Delete(ctx context.Context, player entitiy.Players) (entitiy.Players, error) {
	script := "DELETE FROM players WHERE id = ?"

	result, err := repo.DB.ExecContext(ctx, script, player.Id)
	if err != nil {
		return player, err
	}
	rowCnt, err := result.RowsAffected()
	if rowCnt == 0 {
		return player, err
	}
	return player, err
}
