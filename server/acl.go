package server

import (
	"time"

	"github.com/zicare/rgm/postgres"
)

//ACL struct exported
type ACL struct {
	postgres.Table
	Method string    `db:"method"     json:"method"`
	RoleID string    `db:"role_id"    json:"role"`
	Start  time.Time `db:"start_date" json:"from"`
	End    time.Time `db:"end_date"   json:"to"`
	Route  string    `db:"route"      json:"route"`
}

//Name exported
func (ACL) Name() string {
	return "view_grants"
}
