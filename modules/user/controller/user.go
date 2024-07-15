package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
)

// UserController exported
type UserController struct {
	crud ctrl.CrudController
}

// Fetches a users list
func (uc UserController) Fetch(c *gin.Context) {
	uc.crud.Fetch(c, new(model.User))
}

// Find a user
func (uc UserController) Find(c *gin.Context) {
	uc.crud.Find(c, new(model.User))
}

// Create a new user
func (uc UserController) Post(c *gin.Context) {
	uc.crud.Post(c, new(model.User))
}

// Update an existing user
func (uc UserController) Update(c *gin.Context) {
	uc.crud.Update(c, new(model.User))
}

// Delete a user
func (uc UserController) Delete(c *gin.Context) {
	uc.crud.Delete(c, new(model.User))
}
