package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// #############################################################################
// API

func apiGetRigInfo(c *gin.Context) {
	response := rigConnection.cmdGetRigInfo()
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdGetRigInfo() CommandResponse {
	raw := c.command("+\xf5")

	return CommandResponse{Success: true,
		Raw:  raw,
		Data: ""}
}
