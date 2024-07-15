package agent

import (
	ctrl "gonitor-server/modules/agent/controller"
	"gonitor-server/modules/agent/model"

	"github.com/gin-gonic/gin"
	"github.com/zicare/rgm"
	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/lib"
	"github.com/zicare/rgm/postgres"
)

// Export Routes
func Routes(r *gin.Engine) {

	g := r.Group("a")

	/********* no authentication required *********/

	/*
		story := new(ctrl.StoryController)
		g.GET("/stories", story.Fetch)
		g.HEAD("/stories", story.FetchHead)
		g.GET("/stories/:story_id", story.Find)
	*/

	/********* basic authentication required - rgm.BHC *********/

	fn := ds.UserDSFactory(postgres.UserDSFactory)

	event := new(ctrl.EventController)
	g.GET("/events", rgm.BHC(fn, new(model.Account), lib.Crypto(), event.Fetch)...)
	g.GET("/events/:event_id", rgm.BHC(fn, new(model.Account), lib.Crypto(), event.Find)...)
	g.POST("/events", rgm.BHC(fn, new(model.Account), lib.Crypto(), event.Post)...)

	/********* jwt authentication required - rgm.JHC *********/

	/*
		account := new(ctrl.AccountController)
		g.GET("/account", rgm.JHC(account.Find)...)
	*/

}
