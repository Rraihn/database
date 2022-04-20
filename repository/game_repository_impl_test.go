package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestGameInsert(t *testing.T) {
	gameRepository := NewGameRepositoryImpl(go_database.GetConnection())
	ctx := context.Background()
	games := entitiy.Games{
		Name:  "GTA Chinatown,",
		Genre: "Open World",
	}
	result, err := gameRepository.Insert(ctx, games)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestGameFindById(t *testing.T) {
	gameRepository := NewGameRepositoryImpl(go_database.GetConnection())

	games, err := gameRepository.FindById(context.Background(), 1)
	if err != err {
		panic(err)
	}
	fmt.Println(games)
}

func TestGameFindAll(t *testing.T) {
	gameRepository := NewGameRepositoryImpl(go_database.GetConnection())

	games, err := gameRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, game := range games {
		fmt.Println(game)
	}
}

func TestGameUpdate(t *testing.T) {
	gameRepository := NewGameRepositoryImpl(go_database.GetConnection())

	games, err := gameRepository.Update(context.Background(), entitiy.Games{Name: "Lego Batman", Genre: "Lego RPG", Id: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(games)
}

func TestGameDelete(t *testing.T) {
	gameRepository := NewGameRepositoryImpl(go_database.GetConnection())

	games, err := gameRepository.Delete(context.Background(), entitiy.Games{Id: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(games)
}
