package repository

import (
	"context"
	"go_database/entitiy"
)

type CarRepository interface {
	Insert(ctx context.Context, car entitiy.Cars) (entitiy.Cars, error)
	FindById(ctx context.Context, id int32) (entitiy.Cars, error)
	FindAll(ctx context.Context) ([]entitiy.Cars, error)
	Update(ctx context.Context, car *entitiy.Cars) (*entitiy.Cars, error)
	Delete(ctx context.Context, id int32) (bool, error)
}
