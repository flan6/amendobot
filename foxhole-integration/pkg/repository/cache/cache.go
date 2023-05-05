package cache

import (
	tcache "github.com/jellydator/ttlcache/v3"

	"notifier/pkg/entity"
	"notifier/pkg/repository/cache/ttlcache"
)

type Cache[C entity.Cacheable] interface {
	Save(entity C) error
	Get(key string) (C, error)
	Delete(key string)
}

func NewCache[E entity.Cacheable](cache *tcache.Cache[string, any]) Cache[E] {
	return ttlcache.NewCache[E](cache)
}
