package controller

import (
	"gonitor-server/modules/user/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm/ctrl"
	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/lib"
	"github.com/zicare/rgm/postgres"
)

// PinController exported
type PinController struct {
	base ctrl.PinController
}

// Create pin and sends it back to user
func (pc PinController) Post(c *gin.Context) {

	fn := ds.PinDSFactory(postgres.PinDSFactory)
	pc.base.Post(c, fn, new(model.Pin), new(model.Account))
}

// Update password using pin
func (pc PinController) Patch(c *gin.Context) {

	fn := ds.PinDSFactory(postgres.PinDSFactory)
	pc.base.Patch(c, fn, new(model.Pin), new(model.Account), lib.Crypto())
}
