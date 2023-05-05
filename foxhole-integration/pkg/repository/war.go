package repository

import (
	"context"

	"notifier/pkg/entity"
	"notifier/pkg/repository/cache"
	"notifier/pkg/repository/client"
	"notifier/pkg/repository/storage"
)

type WarRepository interface {
	GetLatestWarState(ctx context.Context) (*entity.War, error)
}

func newWarRepository(
	client client.Client[*entity.War],
	cache cache.Cache[*entity.War],
	storage storage.Storage[*entity.War],
) WarRepository {
	return warRepository{
		client,
		cache,
		storage,
	}
}

type warRepository struct {
	client  client.Client[*entity.War]
	cache   cache.Cache[*entity.War]
	storage storage.Storage[*entity.War]
}

func (w warRepository) GetLatestWarState(ctx context.Context) (*entity.War, error) {
	war, err := w.client.Get()
	if err != nil {
		return nil, err
	}

	storedWar, err := w.storage.Save(ctx, war)
	if err != nil {
		return nil, err
	}

	return storedWar, nil
}
