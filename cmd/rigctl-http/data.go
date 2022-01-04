package main

import "strings"

type simpleResponse interface {
	simpleParse(s string)
}

type simpleListResponse interface {
	simpleListParse(s string)
}

type cmdOptions struct {
	Options []string `json:"options"`
}

func (o *cmdOptions) simpleListParse(s string) {
	o.Options = strings.Split(strings.TrimSpace(s), " ")
}
