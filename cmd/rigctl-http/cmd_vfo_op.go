package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rigOptionsVFOOp cmdOptions

// #############################################################################
// API

func apiVFOOp_options(c *gin.Context) {
	response := rigConnection.cmdVFOOp_options()
	c.IndentedJSON(http.StatusOK, response)
}

func apiExecVFOOp(c *gin.Context) {
	var o VFOOp
	if err := c.BindJSON(&o); err != nil {
		return
	}
	response := rigConnection.cmdExecVFOOp(o)
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// API Request/Response Unmarshal/Marshal types

type VFOOp struct {
	Op string `json:"op"`
}

func (o *VFOOp) simpleParse(s string) {
	o.Op = s
}

// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdExecVFOOp(o VFOOp) CommandResponse {
	// TODO unittest for this!
	if rigOptionsVFOOp.Options == nil {
		c.cmdVFOOp_options()
	}

	if !hasOption(rigOptionsVFOOp.Options, o.Op) {
		return CommandResponse{
			Success: false,
			Error:   fmt.Sprintf("Option not available on this rig, see data for available options"),
			Data:    rigOptionsVFOOp}
	}

	raw := c.command(fmt.Sprintf("+G %s", o.Op))
	return telnetParseSimpleResponse(raw, &o)
}

func (c *TelnetRigConnection) cmdVFOOp_options() CommandResponse {
	raw := c.command("+G ?")

	// sets rigOptionsVFOOP so we have a cache for cmdExec
	return telnetParseSimpleListResponse(raw, &rigOptionsVFOOp)
}
