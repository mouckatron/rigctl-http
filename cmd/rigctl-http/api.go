package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func routerPaths(r *gin.Engine) {
  r.GET("/frequency", getFreq)
  r.POST("/frequency", setFreq)
  r.GET("/powerstat", getPowerstat)
  r.POST("/powerstat", setPowerstat)
  r.GET("/rig_info", getRigInfo)
  r.GET("/dump_caps", getDumpCaps)
}

func getFreq(c *gin.Context) {
  response := rigConnection.GetFreq()
  c.IndentedJSON(http.StatusOK, response)
}

func setFreq(c *gin.Context) {
  var f Frequency
  if err := c.BindJSON(&f); err != nil {
    return
  }
  response := rigConnection.SetFreq(f)
  c.IndentedJSON(http.StatusOK, response)
}

func getPowerstat(c *gin.Context) {
  response := rigConnection.GetPowerstat()
  c.IndentedJSON(http.StatusOK, response)
}

func setPowerstat(c *gin.Context) {
  var p Powerstat
  if err := c.BindJSON(&p); err != nil {
    return
  }
  response := rigConnection.SetPowerstat(p)
  c.IndentedJSON(http.StatusOK, response)
}

func getRigInfo(c *gin.Context) {
  response := rigConnection.GetRigInfo()
  c.IndentedJSON(http.StatusOK, response)
}

func getDumpCaps(c *gin.Context) {
  response := rigConnection.DumpCaps()
  c.IndentedJSON(http.StatusOK, response)
}
