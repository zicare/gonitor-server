package server

import (
	"gonitor-server/modules/agent"
	"gonitor-server/modules/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/zicare/rgm/config"
	"golang.org/x/crypto/acme/autocert"
)

//Init sets the router
func Init(environment string) {

	var (
		c = config.Config()
		r = gin.New()
	)

	//set global middleware
	mw(r)

	//load route groups
	//public.Routes(r)
	agent.Routes(r)
	user.Routes(r)

	if c.GetString("server.domain") == "localhost" {
		r.Run(c.GetString("server.ip") + ":" + c.GetString("server.port"))
	} else {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(c.GetString("server.domain")),
			Cache:      autocert.DirCache("/path/to/autocert/cache/gonitor-server/certs"),
		}
		glog.Fatal(autotls.RunWithManager(r, &m))
	}
}

func mw(r *gin.Engine) {

	//router.Use(gin.Logger())
	//router.Use(ginglog.Logger(3 * time.Second))
	//router.Use(mw.Auth())
	//router.Use(mw.AppKeyCheck())
	//router.Use(mw.BadRequestLogger(new(BadRequestLogger)))

	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	//corsConfig.AllowMethods = []string{"PUT", "PATCH"}
	//corsConfig.AllowCredentials = true
	//corsConfig.MaxAge = 12 * time.Hour
	corsConfig.AllowOriginFunc = func(origin string) bool {
		//AppKeyCheck midleware will take care of this
		return true
		//return origin == "http://localhost:4200" //validate with a db
	}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-App-Key"}
	corsConfig.ExposeHeaders = []string{"X-Range", "X-Checksum", "Date"}

	r.Use(cors.New(corsConfig))

}
