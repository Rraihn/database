package repository_player

import (
	"context"
	"go_database/entitiy"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player entitiy.Player) (entitiy.Player, error)
	FindById(ctx context.Context, id int32) (entitiy.Player, error)
	FindAll(ctx context.Context) ([]entitiy.Player, error)
	Update(ctx context.Context, player entitiy.Player) (entitiy.Player, error)
	Delete(ctx context.Context) (entitiy.Player, error)
}
