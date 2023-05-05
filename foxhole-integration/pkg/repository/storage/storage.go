package storage

import (
	"context"

	"notifier/pkg/entity"
	"notifier/pkg/repository/storage/badgerdb"
)

type Storage[Type entity.Storable] interface {
	Save(ctx context.Context, data Type) (Type, error)
	Get(ctx context.Context, key string) (Type, error)
	GetFiltered(ctx context.Context, filters ...entity.Filter[Type]) ([]Type, error)
	CountFiltered(ctx context.Context, filters ...entity.Filter[Type]) (uint64, error)
	Delete(ctx context.Context, key string) error
	Sequence(ctx context.Context, key string) (uint64, error)
}

func NewStorage[E entity.Storable]() Storage[E] {
	return badgerdb.NewRepository[E]()
}
