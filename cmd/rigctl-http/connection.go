package main

// interfaced in case I am ever able to get hamlib inside Go!
type RigConnection interface {
	init()
	cmdSetFreq(f Frequency) CommandResponse
	cmdGetFreq() CommandResponse
	cmdSetPowerstat(p Powerstat) CommandResponse
	cmdGetPowerstat() CommandResponse
	cmdGetRigInfo() CommandResponse
	cmdDumpCaps() CommandResponse
	cmdExecVFOOp(o VFOOp) CommandResponse
	cmdVFOOp_options() CommandResponse
  cmdGetMode() CommandResponse
  cmdSetMode(m Mode) CommandResponse
  cmdMode_options() CommandResponse
}

//TODO raw should disappear in future when it is further developed
type CommandResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Raw     string      `json:"raw"`
	Data    interface{} `json:"data"`
}
