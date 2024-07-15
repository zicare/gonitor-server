package controller

import (
	"gonitor-server/modules/agent/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
)

// EventController exported
type EventController struct {
	crud ctrl.CrudController
}

// Fetches an events list
func (ec EventController) Fetch(c *gin.Context) {
	ec.crud.Fetch(c, new(model.Event))
}

// Find an event
func (ec EventController) Find(c *gin.Context) {
	ec.crud.Find(c, new(model.Event))
}

// Create a new event
func (ec EventController) Post(c *gin.Context) {
	ec.crud.Post(c, new(model.Event))
}
