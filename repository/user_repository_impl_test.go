package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestUserInsert(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	ctx := context.Background()
	user := entitiy.Users{
		Username: "Rio",
		Password: "Dos-78",
	}
	result, err := userRepository.Insert(ctx, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUserFindById(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	result, err := userRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUserFindAll(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	result, err := userRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUserUpdate(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	result, err := userRepository.Update(context.Background(), entitiy.Users{Username: "Hans", Password: "jojo", Id: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUserDelete(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	result, err := userRepository.Delete(context.Background(), entitiy.Users{Id: 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
