package model

import (
	"database/sql"

	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/postgres"
)

//Endpoint struct exported
type Endpoint struct {
	postgres.Table
	EndpointID *int64 `db:"endpoint_id"   json:"endpoint_id"    binding:"omitempty"  pk:"1"`
	Method     string `db:"method"        json:"method"         binding:"required,oneof=GET HEAD POST PUT PATCH DELETE"`
	Route      string `db:"route"         json:"route"          binding:"required,min=1,max=255"`
}

//Name exported
func (Endpoint) Name() string {
	return "endpoints"
}

//  BeforeInsert exported
func (t *Endpoint) BeforeInsert(qo *ds.QueryOptions, tx *sql.Tx) error {

	return t.BeforeUpdate(qo, tx)
}

//  BeforeUpdate exported
func (t *Endpoint) BeforeUpdate(qo *ds.QueryOptions, tx *sql.Tx) error {
	return nil
}
