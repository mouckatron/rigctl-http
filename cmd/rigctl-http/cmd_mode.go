package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rigOptionsMode cmdOptions

// #############################################################################
// API

func apiMode_options(c *gin.Context) {
  response := rigConnection.cmdMode_options()
  c.IndentedJSON(http.StatusOK, response)
}

func apiGetMode(c *gin.Context) {
	response := rigConnection.cmdGetMode()
	c.IndentedJSON(http.StatusOK, response)
}

func apiSetMode(c *gin.Context) {
	var m Mode
	if err := c.BindJSON(&m); err != nil {
		return
	}
	response := rigConnection.cmdSetMode(m)
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// API Request/Response Unmarshal/Marshal types

type Mode struct {
	Mode string `json:"mode"`
  Passband int `json:"passband"`
}

func (o *Mode) simpleParse(s string) {
}

func (o *Mode) kvParse(m map[string]string) {
  o.Mode = m["Mode"]
  o.Passband = rawToInt(m["Passband"])
}

// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdGetMode() CommandResponse {
  raw := c.command("+m")
	return telnetParseKVResponse(raw, &Mode{})
}

func (c *TelnetRigConnection) cmdSetMode(m Mode) CommandResponse {
  if rigOptionsMode.Options == nil {
    c.cmdMode_options()
  }

  if !hasOption(rigOptionsMode.Options, m.Mode) {
		return CommandResponse{
			Success: false,
			Error:   fmt.Sprintf("Option not available on this rig, see data for available options"),
			Data:    rigOptionsMode}
  }
    
	raw := c.command(fmt.Sprintf("+M %s %d", m.Mode, m.Passband))
	return telnetParseSimpleResponse(raw, &m)
}

func (c *TelnetRigConnection) cmdMode_options() CommandResponse {
  raw := c.command("+M ?")

  // sets rigOptionsMode so we have a cache for cmdGet/cmdSet
  return telnetParseSimpleListResponse(raw, &rigOptionsMode)
}
