package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
)

// RoleController exported
type RoleController struct {
	crud ctrl.CrudController
}

// Fetches a roles list
func (pc RoleController) Fetch(c *gin.Context) {
	pc.crud.Fetch(c, new(model.Role))
}
