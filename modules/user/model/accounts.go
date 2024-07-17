package model

import (
	"database/sql"
	"time"

	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/lib"
	"github.com/zicare/rgm/postgres"
)

//  Account struct exported
type Account struct {
	postgres.Table
	UserID      string    `db:"user_id"        view:"1"   json:"uid"         binding:"required"   pk:"1"`
	RoleID      string    `db:"role_id"        view:"1"   json:"role"`
	TPS         float32   `db:"tps"            view:"1"   json:"tps"`
	Email       string    `db:"email"                     json:"usr"         binding:"required,email"`
	Password    string    `db:"password"                  json:"pwd"         binding:"required,min=8"`
	AccessStart time.Time `db:"access_start"   view:"1"   json:"from"`
	AccessEnd   time.Time `db:"access_end"     view:"1"   json:"to"`
}

//  Name exported
func (Account) Name() string {
	return "users"
}

//  BeforeUpdate exported
func (t *Account) BeforeUpdate(qo *ds.QueryOptions, tx *sql.Tx) error {

	t.Password = lib.Crypto().Encode(t.Password)
	return nil
}
