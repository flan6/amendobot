package repository

import (
	"context"
	"fmt"

	"notifier/pkg/entity"
	"notifier/pkg/repository/cache"
	"notifier/pkg/repository/client"
	"notifier/pkg/repository/storage"
)

type MapsRepository interface {
	GetMaps(ctx context.Context) (*entity.Maps, error)
}

func newMapsRepository(
	client client.Client[*entity.Maps],
	cache cache.Cache[*entity.Maps],
	storage storage.Storage[*entity.Maps],
) MapsRepository {
	return mapsRepository{
		client,
		cache,
		storage,
	}
}

type mapsRepository struct {
	client  client.Client[*entity.Maps]
	cache   cache.Cache[*entity.Maps]
	storage storage.Storage[*entity.Maps]
}

func (m mapsRepository) GetMaps(ctx context.Context) (*entity.Maps, error) {
	maps, err := m.cache.Get("maps")
	if err == nil {
		return maps, nil
	}

	maps, err = m.storage.Get(ctx, maps.Key())
	if err != nil {
		// TODO : logger
		fmt.Printf("error in maps.storage.Get: %s\n", err.Error())
		// return nil, err
	}

	if maps != nil {
		return maps, nil
	}

	maps, err = m.client.Get()
	if err != nil {
		return nil, err
	}

	_, err = m.storage.Save(ctx, maps)
	if err != nil {
		return nil, err
	}

	return maps, nil
}
