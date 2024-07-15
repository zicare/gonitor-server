package model

import (
	"time"

	"github.com/zicare/rgm/postgres"
)

//Event struct exported
type Event struct {
	postgres.Table
	EventID  *int64    `db:"event_id"          json:"event_id"          binding:"omitempty" pk:"1"`
	HostID   int64     `db:"host_id"           json:"host_id"           binding:"required"`
	ItemID   int64     `db:"item_id"           json:"item_id"           binding:"required"`
	Data     float64   `db:"data"              json:"data"              binding:"exists"`
	Datetime time.Time `db:"datetime"          json:"datetime"          binding:"required"`
}

//Name exported
func (Event) Name() string {
	return "events"
}
