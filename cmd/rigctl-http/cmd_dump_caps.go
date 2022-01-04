package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// #############################################################################
// API

func apiDumpCaps(c *gin.Context) {
	response := rigConnection.cmdDumpCaps()
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdDumpCaps() CommandResponse {
	raw := c.command("+1")
	return CommandResponse{Success: true,
		Raw:  raw,
		Data: ""}
}
