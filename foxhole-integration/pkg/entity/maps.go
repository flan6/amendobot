package entity

import "time"

const (
	mapsEndpoint = "worldconquest/maps"
	mapsId       = "maps"
)

type Maps []Map

type Map string

func (m *Maps) Id() string {
	return mapsId
}

func (m *Maps) SetId(id string) {}

func (m *Maps) GenId() {}

func (m *Maps) GenCacheKey() string {
	return mapsId
}

func (m *Maps) Ttl() time.Duration {
	return time.Hour * 24
}

func (m *Maps) Name() string {
	return Name(m)
}

func (m *Maps) Key() string {
	return Key(m)
}

func (m *Maps) ApiEndpoint() string {
	return mapsEndpoint
}

func (m *Maps) Etag() string {
	return ""
}

func (m *Maps) SetEtag(_ string) {}
