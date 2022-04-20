package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestItemInsert(t *testing.T) {
	itemRepository := NewItemRepository(go_database.GetConnection())
	ctx := context.Background()
	items := entitiy.Items{
		Name: "Phone",
		Qty:  15,
	}
	result, err := itemRepository.Insert(ctx, items)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestItemFindById(t *testing.T) {
	itemRepository := NewItemRepository(go_database.GetConnection())

	items, err := itemRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}

func TestItemFindAll(t *testing.T) {
	itemRepository := NewItemRepository(go_database.GetConnection())

	items, err := itemRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}

func TestItemUpdate(t *testing.T) {
	ItemRepository := NewItemRepository(go_database.GetConnection())

	items, err := ItemRepository.Update(context.Background(), entitiy.Items{Name: "Sosis", Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}

func TestItemDelete(t *testing.T) {
	ItemRepository := NewItemRepository(go_database.GetConnection())

	items, err := ItemRepository.Delete(context.Background(), entitiy.Items{Id: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}
