package user

import (
	ctrl "gonitor-server/modules/user/controller"
	"gonitor-server/modules/user/model"

	bctrl "github.com/zicare/rgm/ctrl"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm"
	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/lib"
	"github.com/zicare/rgm/postgres"
)

// Export Routes
func Routes(r *gin.Engine) {

	g := r.Group("u")

	/********* no authentication required *********/

	pin := new(ctrl.PinController)
	g.POST("/pin", pin.Post)   // create pin and sends it back to user
	g.PATCH("/pwd", pin.Patch) // updates password using pin

	/********* basic authentication required - rgm.BHC *********/

	jwt := new(bctrl.JwtController)
	fn := ds.UserDSFactory(postgres.UserDSFactory)
	g.GET("/jwt", rgm.BHC(fn, new(model.Account), lib.Crypto(), jwt.Get)...) // get a JWT

	/********* jwt authentication required - rgm.JHC *********/

	account := new(ctrl.AccountController)
	g.GET("/account", rgm.JHC(account.Find)...)

	role := new(ctrl.RoleController)
	g.GET("/roles", rgm.JHC(role.Fetch)...)

	endpoint := new(ctrl.EndpointController)
	g.GET("/endpoints", rgm.JHC(endpoint.Fetch)...)

	user := new(ctrl.UserController)
	g.GET("/users", rgm.JHC(user.Fetch)...)
	g.GET("/users/:user_id", rgm.JHC(user.Find)...)
	g.POST("/users", rgm.JHC(user.Post)...)
	g.PUT("/users/:user_id", rgm.JHC(user.Update)...)
	g.DELETE("/users/:user_id", rgm.JHC(user.Delete)...)

	event := new(ctrl.EventController)
	g.GET("/events", rgm.JHC(event.Fetch)...)

}
