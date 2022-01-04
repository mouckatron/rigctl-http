package main

import (
  "github.com/gin-gonic/gin"
)

var rigConnection RigConnection

func main() {

	// get arguments

	// decide if we're using hamlib or telnet (always telnet for now)

	// start a RigConnection interface that implements the functions
	rigConnection = &TelnetRigConnection{Address: "localhost:4532"}

	rigConnection.init()

	// start the API server which uses the RigConnection thing
  router := gin.Default()
  routerPaths(router)
  router.Run("localhost:8080")
}
