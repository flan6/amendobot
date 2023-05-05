package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	mapDataEndpointStatic  = "maps/%s/static"
	mapDataEndpointDynamic = "maps/%s/dynamic/public"
)

type MapData struct {
	ID                   string
	Map                  Map
	Dynamic              bool
	RegionID             int           `json:"regionId"`
	ScorchedVictoryTowns int           `json:"scorchedVictoryTowns"`
	MapItems             []MapItem     `json:"mapItems"`
	MapTextItems         []MapTextItem `json:"mapTextItems"`
	LastUpdated          int64         `json:"lastUpdated"`
	Version              int           `json:"version"`
}

func (m *MapData) Id() string {
	return m.ID
}

func (m *MapData) SetId(id string) {
	m.ID = id
}

func (m *MapData) GenId() {
	m.ID = GenId(m)
}

func (m *MapData) GenCacheKey() string {
	return Key(m)
}

func (m *MapData) Ttl() time.Duration {
	return time.Hour
}

func (m *MapData) Name() string {
	return Name(m)
}

func (m *MapData) Key() string {
	return Key(m)
}

func (m *MapData) ApiEndpoint() string {
	if m.Dynamic {
		return fmt.Sprintf(mapDataEndpointDynamic, m.Map)
	}

	return fmt.Sprintf(mapDataEndpointStatic, m.Map)
}

func (m *MapData) Etag() string {
	return strconv.Itoa(m.Version)
}

func (m *MapData) SetEtag(e string) {
	str := strings.Trim(e, `"`)
	tag, err := strconv.Atoi(str)
	if err != nil {
		// TODO : log somehow
		panic(err)
	}

	m.Version = tag
}

type MapItem struct {
	TeamID   string  `json:"teamId"`
	IconType int     `json:"iconType"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Flags    int     `json:"flags"`
}

type MapTextItem struct {
	Text          string        `json:"text"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	MapMarkerType MapMarkerType `json:"mapMarkerType"`
}

type MapMarkerType string

const (
	MapMarkerMajor MapMarkerType = "Major"
	MapMarkerMinor MapMarkerType = "Minor"
)
