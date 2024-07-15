package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/lib"
	"github.com/zicare/rgm/postgres"
)

//User struct exported
type User struct {
	postgres.Table
	UserID      *int64    `db:"user_id"            json:"user_id"              binding:"omitempty" pk:"1"`
	RoleID      *int64    `db:"role_id"            json:"role_id"              binding:"omitempty"`
	TPS         float32   `db:"tps"                json:"-"                    binding:"-"`
	Email       string    `db:"email"              json:"email"                binding:"required,email"`
	Password    *string   `db:"password"           json:"-"                    binding:"-"`
	AccessStart time.Time `db:"access_start"       json:"access_start"         binding:"required"`
	AccessEnd   time.Time `db:"access_end"         json:"access_end"           binding:"required"`
	Role        *Role     `db:"-"                  json:"role,omitempty"`
}

//Name exported
func (User) Name() string {
	return "users"
}

//  BeforeInsert exported
func (t *User) BeforeInsert(qo *ds.QueryOptions, tx *sql.Tx) error {

	return t.BeforeUpdate(qo, tx)
}

//  BeforeUpdate exported
func (t *User) BeforeUpdate(qo *ds.QueryOptions, tx *sql.Tx) error {

	t.TPS = 10

	return nil
}

//  AfterSelect exported
func (t *User) AfterSelect(qo *ds.QueryOptions) error {

	if lib.Contains(qo.Dig, "users.role") && t.RoleID != nil {
		t.Role = new(Role)
		p := make(ds.Params)
		p["role_id"] = fmt.Sprint(*t.RoleID)
		if _, _, err := t.Find(qo.Copy(t.Role, p)); err != nil {
			return err
		}
	}

	return nil
}
