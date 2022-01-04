package main

import (
	"flag"
	"fmt"

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

	// get arguments
	flag.StringVar(&myHost, "host", "localhost", myHostHelp)
	flag.IntVar(&myPort, "port", 8080, myPortHelp)

	flag.StringVar(&rcHost, "rchost", "localhost", rcHostHelp)
	flag.IntVar(&rcPort, "rcport", 4532, rcPortHelp)

	flag.Parse()

	// decide if we're using hamlib or telnet (always telnet for now)

	// start a RigConnection interface that implements the functions
	rigConnection = &TelnetRigConnection{Address: fmt.Sprintf("%s:%d", rcHost, rcPort)}

	rigConnection.init()

	// start the API server which uses the RigConnection thing
	router := gin.Default()
	routerPaths(router)
	router.Run(fmt.Sprintf("%s:%d", myHost, myPort))
}
