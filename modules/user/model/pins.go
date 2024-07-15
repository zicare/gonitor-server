package model

import (
	"time"

	"github.com/zicare/rgm/postgres"
)

//Partner struct exported
type Pin struct {
	postgres.Table
	PinID      *int64    `db:"pin_id"       json:"-"               pk:"1"`
	Email      string    `db:"email"        json:"email"           binding:"required,email"`
	PIN        string    `db:"pin"          json:"code"`
	Created    time.Time `db:"created"      json:"created"`
	Expiration time.Time `db:"expiration"   json:"expiration"`
}

//Name exported
func (Pin) Name() string {
	return "pins"
}
