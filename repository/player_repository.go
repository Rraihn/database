package repository

import (
	"context"
	"go_database/entitiy"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player entitiy.Players) (entitiy.Players, error)
	FindById(ctx context.Context, id int32) (entitiy.Players, error)
	FindAll(ctx context.Context) ([]entitiy.Players, error)
	Update(ctx context.Context, player *entitiy.Players) (*entitiy.Players, error)
	Delete(ctx context.Context, id int32) (bool, error)
}
