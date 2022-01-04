package main

type RigConnection interface {
  init()
	SetFreq(f Frequency) CommandResponse
	GetFreq() CommandResponse
  SetPowerstat(p Powerstat) CommandResponse
  GetPowerstat() CommandResponse
  GetRigInfo() CommandResponse
  DumpCaps() CommandResponse
}

type CommandResponse struct {
	Success bool `json:"success"`
	Raw     string `json:"raw"`
  Data    interface{} `json:"data"`
}


