package repository

import (
	"context"
	"go_database/entitiy"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entitiy.Comments) (entitiy.Comments, error)
	FindById(ctx context.Context, id int32) (entitiy.Comments, error)
	FindAll(ctx context.Context) ([]entitiy.Comments, error)
}
