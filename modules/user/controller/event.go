package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
)

// EventController exported
type EventController struct {
	crud ctrl.CrudController
}

// Fetches an events list
func (pc EventController) Fetch(c *gin.Context) {
	pc.crud.Fetch(c, new(model.Event))
}
