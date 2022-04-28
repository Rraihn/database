package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestCarInsert(t *testing.T) {
	carRepository := NewCarRepository(go_database.GetConnection())
	ctx := context.Background()
	car := entitiy.Cars{
		Name: "Discovery Evoque",
		Merk: "Range Rover",
	}
	result, err := carRepository.Insert(ctx, car)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCarFindById(t *testing.T) {
	carRepository := NewCarRepository(go_database.GetConnection())

	car, err := carRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(car)
}

func TestCarFindAll(t *testing.T) {
	carRepository := NewCarRepository(go_database.GetConnection())

	cars, err := carRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, car := range cars {
		fmt.Println(car)
	}
}

func TestCarUpdate(t *testing.T) {
	carRepository := NewCarRepository(go_database.GetConnection())
	car, err := carRepository.Update(context.Background(), entitiy.Cars{Name: "unicorn special Edition", Merk: "Hoonigan", Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(car)
}

func TestCarDelete(t *testing.T) {
	carRepository := NewCarRepository(go_database.GetConnection())
	car, err := carRepository.Delete(context.Background(), entitiy.Cars{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(car)
}
