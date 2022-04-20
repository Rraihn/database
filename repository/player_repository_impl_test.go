package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestInsertPlayer(t *testing.T) {
	insertPlayer := NewPlayerRepository(go_database.GetConnection())
	ctx := context.Background()
	players := entitiy.Players{
		Name:     "Aqias",
		Nickname: "Beast",
		Gender:   "boy",
	}
	result, err := insertPlayer.Insert(ctx, players)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func TestPlayerFindById(t *testing.T) {
	playerRepository := NewPlayerRepository(go_database.GetConnection())

	result, err := playerRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestPlayerFindAll(t *testing.T) {
	playerRepository := NewPlayerRepository(go_database.GetConnection())

	result, err := playerRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestPlayerUpdate(t *testing.T) {
	playerRepository := NewPlayerRepository(go_database.GetConnection())

	result, err := playerRepository.Update(context.Background(), entitiy.Players{Name: "Raihan", Nickname: "tequila", Gender: "boys", Id: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestPlayerDelete(t *testing.T) {
	playerRepository := NewPlayerRepository(go_database.GetConnection())

	result, err := playerRepository.Delete(context.Background(), entitiy.Players{Id: 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
