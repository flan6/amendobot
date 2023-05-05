package entity

import (
	"time"
)

const warEndpoint = "worldconquest/war"

type War struct {
	ID                   string `json:"id"`
	WarID                string `json:"warId"`
	WarNumber            int    `json:"warNumber"`
	Winner               Winner `json:"winner"`
	ConquestStartTime    int    `json:"conquestStartTime"`
	ConquestEndTime      int    `json:"conquestEndTime"`
	ResistanceStartTime  int    `json:"resistanceStartTime"`
	RequiredVictoryTowns int    `json:"requiredVictoryTowns"`
}

type Winner string

const (
	WinnerNone      Winner = "NONE"
	WinnerColonials Winner = "COLONIALS"
	WinnerWardens   Winner = "WARDENS"
)

func (w *War) Id() string {
	if w.ID == "" {
		w.GenId()
	}

	return w.ID

}

func (w *War) SetId(id string) {
	w.ID = id
}

func (w *War) GenId() {
	w.ID = GenId(w)
}

func (w *War) GenCacheKey() string {
	return Key(w)
}

func (w *War) Ttl() time.Duration {
	return time.Hour
}

func (w *War) ApiEndpoint() string {
	return warEndpoint
}

// War is not cacheable by the api. So no etags.
func (w *War) Etag() string {
	return ""
}

// War is not cacheable by the api. So no etags.
func (w *War) SetEtag(_ string) {}

func (w War) Name() string {
	return Name(w)
}

func (w War) Key() string {
	return Key(&w)
}
