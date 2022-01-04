package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// #############################################################################
// API

func apiGetFreq(c *gin.Context) {
	response := rigConnection.cmdGetFreq()
	c.IndentedJSON(http.StatusOK, response)
}

func apiSetFreq(c *gin.Context) {
	var f Frequency
	if err := c.BindJSON(&f); err != nil {
		return
	}
	response := rigConnection.cmdSetFreq(f)
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// API Request/Response Unmarshal/Marshal types

type Frequency struct {
	Frequency string `json:"frequency"`
}

func (o *Frequency) simpleParse(s string) {
	o.Frequency = s
}

// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdGetFreq() CommandResponse {
	raw := c.command("+f")
	return telnetParseSimpleResponse(raw, &Frequency{})
}

func (c *TelnetRigConnection) cmdSetFreq(f Frequency) CommandResponse {
	raw := c.command(fmt.Sprintf("+F %s", f.Frequency))
	return telnetParseSimpleResponse(raw, &f)
}
