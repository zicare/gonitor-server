package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
)

// Endpointontroller exported
type EndpointController struct {
	crud ctrl.CrudController
}

// Fetches an endpoints list
func (pc EndpointController) Fetch(c *gin.Context) {
	pc.crud.Fetch(c, new(model.Endpoint))
}
