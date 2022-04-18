package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entitiy"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entitiy.Comments{
		Email:   "raihan@gmail.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
