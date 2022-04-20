package repository

import (
	"context"
	"go_database/entitiy"
)

type ItemRepository interface {
	Insert(ctx context.Context, item entitiy.Items) (entitiy.Items, error)
	FindById(ctx context.Context, id int32) (entitiy.Items, error)
	FindAll(ctx context.Context) ([]entitiy.Items, error)
	Update(ctx context.Context, item entitiy.Items) (entitiy.Items, error)
	Delete(ctx context.Context, items entitiy.Items) (entitiy.Items, error)
}
