package main

import (
	"github.com/gin-gonic/gin"
)

func routerPaths(r *gin.Engine) {
	r.GET("/frequency", apiGetFreq)
	r.POST("/frequency", apiSetFreq)
	r.GET("/powerstat", apiGetPowerstat)
	r.POST("/powerstat", apiSetPowerstat)
	r.GET("/rig_info", apiGetRigInfo)
	r.GET("/dump_caps", apiDumpCaps)
	r.GET("/vfo_op/_opts", apiVFOOp_options)
	r.POST("/vfo_op", apiExecVFOOp)
}
