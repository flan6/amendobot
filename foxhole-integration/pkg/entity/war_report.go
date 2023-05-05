package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const WarReportEndpoint = "worldconquest/warReport/%s"

type WarReport struct {
	ID                 string
	Map                Map
	TotalEnlistments   int `json:"totalEnlistments"`
	ColonialCasualties int `json:"colonialCasualties"`
	WardenCasualties   int `json:"wardenCasualties"`
	DayOfWar           int `json:"dayOfWar"`
	Version            int `json:"version"`
}

func (m *WarReport) Id() string {
	if m.ID == "" {
		m.GenId()
	}

	return m.ID
}

func (m *WarReport) SetId(id string) {
	m.ID = id
}

func (m *WarReport) GenId() {
	m.ID = GenId(m)
}

func (m *WarReport) GenCacheKey() string {
	return Key(m)
}

func (m *WarReport) Ttl() time.Duration {
	return time.Hour
}

func (m *WarReport) Name() string {
	return Name(m)
}

func (m *WarReport) Key() string {
	return Key(m)
}

func (m *WarReport) ApiEndpoint() string {
	return fmt.Sprintf(WarReportEndpoint, m.Map)
}

func (m *WarReport) Etag() string {
	return strconv.Itoa(m.Version)
}

func (m *WarReport) SetEtag(e string) {
	str := strings.Trim(e, `"`)
	tag, err := strconv.Atoi(str)
	if err != nil {
		// TODO : log somehow
		panic(err)
	}

	m.Version = tag
}
