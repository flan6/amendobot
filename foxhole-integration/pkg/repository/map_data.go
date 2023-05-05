package repository

import (
	"notifier/pkg/entity"
	"notifier/pkg/repository/cache"
	"notifier/pkg/repository/client"
	"notifier/pkg/repository/storage"
)

type MapDataRepository interface{}

func newMapDataRepository(
	client client.Client[*entity.MapData],
	cache cache.Cache[*entity.MapData],
	storage storage.Storage[*entity.MapData],
) MapDataRepository {
	return mapDataRepository{
		client,
		cache,
		storage,
	}
}

type mapDataRepository struct {
	client  client.Client[*entity.MapData]
	cache   cache.Cache[*entity.MapData]
	storage storage.Storage[*entity.MapData]
}
