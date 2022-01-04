package main

type simpleResponse interface {
  simpleParse(s string)
}

type Frequency struct {
  Frequency string `json:"frequency"`
}

func (o Frequency) simpleParse(s string) {
  o.Frequency = s
}

type Powerstat struct {
  Status int `json:"status"`
}

func (o Powerstat) simpleParse(s string) {
  o.Status = rawToInt(s)
}
