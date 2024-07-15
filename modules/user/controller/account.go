package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
	"github.com/zicare/rgm/ds"
)

// AccountController exported
type AccountController struct {
	crud ctrl.CrudController
}

// Find my account
func (ac AccountController) Find(c *gin.Context) {
	c.AddParam("user_id", ds.UID(c))
	ac.crud.Find(c, new(model.User))
}

// Update my account
func (ac AccountController) Update(c *gin.Context) {
	c.AddParam("user_id", ds.UID(c))
	ac.crud.Update(c, new(model.Account))
}
