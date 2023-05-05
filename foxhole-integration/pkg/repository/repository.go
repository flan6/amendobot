package repository

import (
	"net/http"

	tcache "github.com/jellydator/ttlcache/v3"

	"notifier/pkg/entity"
	"notifier/pkg/repository/cache"
	"notifier/pkg/repository/client"
	"notifier/pkg/repository/storage"
)

type Repository interface {
	entity.Cacheable
	entity.Storable
	entity.Requestable
}

type All struct {
	WarRepository
	MapsRepository
	MapDataRepository
	WarReportRepository
}

func GetAll(ttlCache *tcache.Cache[string, any], httpClient *http.Client) All {
	return All{
		newWarRepository(
			initDependencies[*entity.War](
				ttlCache,
				httpClient,
			),
		),
		newMapsRepository(
			initDependencies[*entity.Maps](
				ttlCache,
				httpClient,
			),
		),
		newMapDataRepository(
			initDependencies[*entity.MapData](
				ttlCache,
				httpClient,
			),
		),
		newWarReportRepository(
			initDependencies[*entity.WarReport](
				ttlCache,
				httpClient,
			),
		),
	}
}

func initDependencies[R Repository](
	ttlCache *tcache.Cache[string, any],
	httpClient *http.Client,
) (
	client.Client[R],
	cache.Cache[R],
	storage.Storage[R],
) {
	client := client.NewClient[R](httpClient)
	cache := cache.NewCache[R](ttlCache)
	storage := storage.NewStorage[R]()

	return client, cache, storage
}
