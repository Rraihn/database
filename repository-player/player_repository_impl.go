package repository_player

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type playersRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) PlayerRepository {
	return &playersRepositoryImpl{DB: db}
}

func (repo *playersRepositoryImpl) Insert(ctx context.Context, player entitiy.Player) (entitiy.Player, error) {
	script := "INSERT INTO comments(name, nickname, gender) VALUES (?, ?)"
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

func (repo *playersRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Player, error) {
	script := "SELECT id, name, nickname, gender FROM  WHERE id = ? LIMIT 3"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	player := entitiy.Player{}

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

func (repo *playersRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Player, error) {
	script := "SELECT id, name, nickname, gender FROM players"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []entitiy.Player
	for rows.Next() {
		player := entitiy.Player{}
		rows.Scan(&player.Id, &player.Name, &player.Nickname, &player.Gender)
		players = append(players, player)
	}
	return players, nil
}

func (repo *playersRepositoryImpl) Update(ctx context.Context, player entitiy.Player) (entitiy.Player, error) {
	script := "UPDATE players SET Id = ? ,Name= ?,Nickname= ? ,Gender = ?"
	result, err := repo.DB.ExecContext(ctx, script, player.Id, player.Name, player.Nickname, player.Gender)
	if err != nil {
		return player, err
	}
	id, err := result.
}

func (repo *playersRepositoryImpl) Delete(ctx context.Context) (entitiy.Player, error) {
	script := "DELETE FROM player WHERE id = ? LIMIT 3"
	res, err := repo.DB.ExecContext(ctx, script)
	player := entitiy.Player{}
	
	if err != nil {
		return  player, err
	}
}
