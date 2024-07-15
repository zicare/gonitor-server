package main

import (
	"flag"
	"fmt"
	"os"

	"gonitor-server/server"

	"github.com/zicare/rgm"
	"github.com/zicare/rgm/ds"
	"github.com/zicare/rgm/postgres"

	"github.com/gin-gonic/gin"
)

var usage = `
Usage:
	gonitor-server [options]
Options:
	-e <environment>
	Sets the running environment. Set this option to 'test' or 'production', default is 'production'.
	Matching configuration files for each environment must be present at './config' directory. 
	--disable-agent
	Have the built-in monitoring agent disabled on start.
	--verbose
	Outputs initialization messages.
Examples:
	# run gonitor-server in test mode with built-in monitoring agent disabled on start
	gonitor-server -e test --disable-agent
	# run gonitor-server in production mode and activate built-in monitoring agent on start
	gonitor-server
`

func main() {

	opts := rgm.InitOpts{
		Environment:  flag.String("e", "production", ""),
		DisableAgent: flag.Bool("disable-agent", false, ""),
		Verbose:      flag.Bool("verbose", false, ""),
		Messages:     nil,
		AclDSFactory: ds.AclDSFactory(postgres.AclDSFactory),
		Acl:          new(server.ACL),
	}

	flag.Usage = func() {
		fmt.Println(usage)
		os.Exit(1)
	}
	flag.Parse()

	if *opts.Environment != "production" {
		flag.Set("alsologtostderr", "true")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := rgm.Init(opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server.Init(*opts.Environment)
}
