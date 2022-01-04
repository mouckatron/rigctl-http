package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// #############################################################################
// API

func apiGetPowerstat(c *gin.Context) {
	response := rigConnection.cmdGetPowerstat()
	c.IndentedJSON(http.StatusOK, response)
}

func apiSetPowerstat(c *gin.Context) {
	var p Powerstat
	if err := c.BindJSON(&p); err != nil {
		return
	}
	response := rigConnection.cmdSetPowerstat(p)
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// API Request/Response Unmarshal/Marshal types

type Powerstat struct {
	Status int `json:"status"`
}

func (o *Powerstat) simpleParse(s string) {
	o.Status = rawToInt(s)
}

// #############################################################################
// TelnetRigConnection
func (c *TelnetRigConnection) cmdGetPowerstat() CommandResponse {
	raw := c.command("+\x88")
	return telnetParseSimpleResponse(raw, &Powerstat{})
}

func (c *TelnetRigConnection) cmdSetPowerstat(p Powerstat) CommandResponse {
	raw := c.command(fmt.Sprintf("+\x87 %d", p.Status))
	return telnetParseSimpleResponse(raw, &p)
}
