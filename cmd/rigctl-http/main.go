package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var rigConnection RigConnection

func main() {

	var myHost string
	const myHostHelp = "Address rigctl-http will be available at"
	var myPort int
	const myPortHelp = "Port rigctl-http will be available at"
	var rcHost string
	const rcHostHelp = "Address rigctld is running at"
	var rcPort int
	const rcPortHelp = "Port rigctld is running on"
	var ginProduction bool = true

	// normal arguments
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.StringVar(&myHost, "host", "localhost", myHostHelp)
	flag.IntVar(&myPort, "port", 8080, myPortHelp)

	flag.StringVar(&rcHost, "rchost", "localhost", rcHostHelp)
	flag.IntVar(&rcPort, "rcport", 4532, rcPortHelp)

	flag.Parse()

	// hidden args!
	for _, arg := range os.Args[1:] {
		if arg == "-gin-debug" {
			ginProduction = false
		}
	}

	// decide if we're using hamlib or telnet (always telnet for now)

	// start a RigConnection interface that implements the functions
	rigConnection = &TelnetRigConnection{Address: fmt.Sprintf("%s:%d", rcHost, rcPort)}

	rigConnection.init()

	// start the API server which uses the RigConnection thing
	if ginProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	routerPaths(router)
	router.Run(fmt.Sprintf("%s:%d", myHost, myPort))
}
