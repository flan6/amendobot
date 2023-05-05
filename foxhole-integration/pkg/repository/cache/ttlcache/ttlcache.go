package ttlcache

import (
	"encoding/json"
	"errors"

	"github.com/jellydator/ttlcache/v3"

	"notifier/pkg/entity"
)

type ttlCache[E entity.Cacheable] struct {
	cache *ttlcache.Cache[string, any]
}

func NewCache[E entity.Cacheable](c *ttlcache.Cache[string, any]) ttlCache[E] {
	return ttlCache[E]{c}
}

func (c ttlCache[E]) Save(entity E) error {
	rawData, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	c.cache.Set(entity.GenCacheKey(), rawData, entity.Ttl())

	return nil
}

func (c ttlCache[E]) Get(key string) (E, error) {
	var i E
	item := c.cache.Get(key)
	if item != nil {
		data := item.Value()
		if err := json.Unmarshal(data.([]byte), &i); err != nil {
			return i, err
		}

		return i, nil
	}

	return i, errors.New("item not found on cache")
}

func (c ttlCache[E]) Delete(key string) {
	c.cache.Delete(key)
}
