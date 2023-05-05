package entity

import (
	"fmt"
	"time"
)

// Entity defines the basic methodos all implementations must follow
type Entity interface {
	// Id must always follow the patern name:unixNano time in nanoseconds.
	// Using unix time it is possible to split and the last part is
	// always the last time this entity was updated.
	Id() string

	SetId(id string)

	// GenID must generate a new id and set it as the new entity ID.
	// A new ID must be generated everytime the entity change value.
	GenId()
}

type Cacheable interface {
	Entity

	GenCacheKey() string
	Ttl() time.Duration
}

type Storable interface {
	Entity

	Name() string
	Key() string
}

type Filter[E Storable] func(t E) bool

// TODO : apply solid principle. Not every requestable has etags.
type Requestable interface {
	Entity

	ApiEndpoint() string
	Etag() string
	SetEtag(etag string)
}

// This way all entities follows same pattern of id generation
// Using unix time it is possible to split and the last part is
// always the last time this entity was updated.
// A new ID must be generated everytime the entity change value.
func GenId[E Entity](e E) string {
	return fmt.Sprintf("%s:%d", Name(e), time.Now().UnixNano())
}

func Name[E any](e E) string {
	return fmt.Sprintf("%T", e)
}

// This way all entities follows same pattern of key generation
func Key[E Entity](e E) string {
	return fmt.Sprintf("%T:%s", e, e.Id())
}
