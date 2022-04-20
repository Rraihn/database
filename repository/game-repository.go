package repository

import (
	"context"
	"go_database/entitiy"
)

type GamesRepository interface {
	Insert(ctx context.Context, game entitiy.Games) (entitiy.Games, error)
	FindById(ctx context.Context, id int32) (entitiy.Games, error)
	FindAll(ctx context.Context) ([]entitiy.Games, error)
	Update(ctx context.Context, game entitiy.Games) (entitiy.Games, error)
	Delete(ctx context.Context, games entitiy.Games) (entitiy.Games, error)
}
