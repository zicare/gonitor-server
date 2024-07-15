package model

import (
	"database/sql"

	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/msg"
	"github.com/zicare/rgm/postgres"
)

//Role struct exported
type Role struct {
	postgres.Table
	RoleID   *int64  `db:"role_id"            json:"role_id"             binding:"omitempty"  pk:"1"`
	Role     string  `db:"role"               json:"role"                binding:"required,min=2,max=64,unik"`
	TPS      float32 `db:"tps"                json:"tps"                 binding:"-"`
	Comments string  `db:"comments"           json:"comments"            binding:"required"`
}

//Name exported
func (Role) Name() string {
	return "roles"
}

//  BeforeInsert exported
func (t *Role) BeforeInsert(qo *ds.QueryOptions, tx *sql.Tx) error {

	return t.BeforeUpdate(qo, tx)
}

//  BeforeUpdate exported
func (t *Role) BeforeUpdate(qo *ds.QueryOptions, tx *sql.Tx) error {

	if t.RoleID != nil && *t.RoleID <= int64(16) {
		//return msg.Get("16").M2E()
		//return new(ds.NotAllowedError)
		err := ds.ValidationErrors{}
		err = append(err, msg.Get("19"))
		return &err
	}
	t.TPS = 10

	return nil
}
