package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// #############################################################################
// API

func apiGetTuningStep(c *gin.Context) {
	response := rigConnection.cmdGetTuningStep()
	c.IndentedJSON(http.StatusOK, response)
}

func apiSetTuningStep(c *gin.Context) {
	var ts TuningStep
	if err := c.BindJSON(&ts); err != nil {
		return
	}
	response := rigConnection.cmdSetTuningStep(ts)
	c.IndentedJSON(http.StatusOK, response)
}

// #############################################################################
// API Request/Response Unmarshal/Marshal types

type TuningStep struct {
	Step int `json:"step"`
}

func (ts *TuningStep) simpleParse(s string) {
	ts.Step = strToInt(s)
}

//
// #############################################################################
// TelnetRigConnection

func (c *TelnetRigConnection) cmdGetTuningStep() CommandResponse {
	raw := c.command("+n")
	return telnetParseSimpleResponse(raw, &TuningStep{})
}

func (c *TelnetRigConnection) cmdSetTuningStep(ts TuningStep) CommandResponse {
	raw := c.command(fmt.Sprintf("+N %d", ts.Step))
	return telnetParseSimpleResponse(raw, &ts)
}
