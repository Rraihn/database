package repository

import (
	"context"
	"go_database/entitiy"
)

type UserRepository interface {
	Insert(ctx context.Context, user entitiy.Users) (entitiy.Users, error)
	FindById(ctx context.Context, id int32) (entitiy.Users, error)
	FindAll(ctx context.Context) ([]entitiy.Users, error)
	Update(ctx context.Context, user entitiy.Users) (entitiy.Users, error)
	Delete(ctx context.Context, user entitiy.Users) (entitiy.Users, error)
}
