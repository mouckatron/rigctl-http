package main

import (
	"bytes"
	"net"
	"strconv"
	"strings"
)

type TelnetRigConnection struct {
	Address string
}

func (c *TelnetRigConnection) init() {}

func (c *TelnetRigConnection) command(cmd string) string {
	conn, err := net.Dial("tcp", c.Address)
	defer conn.Close()
	if err != nil {
		// handle error
	}

	conn.Write([]byte(cmd + "\n"))

	var buf [1024]byte
	response := bytes.NewBuffer(nil)
	last := false

	for {
		n, _ := conn.Read(buf[0:])
		response.Write(buf[0:n])
		length := response.Len()
		contents := response.Bytes()

		if string(contents[length-1:]) == "\n" { // check for the end
			last = true // and mark we thing it is the end
		} else { // if we read more and it does not end \n, keep going
			last = false
		}

		//TODO: Can we use the `RPRT: [0-9]\n` ending to check instead?

		if last {
			break
		}
	}

	return string(response.Bytes())
}

func telnetParseSimpleResponse(r string, o simpleResponse) CommandResponse {
	split := strings.Split(r, "\n")

	var dataline string
	var result string

	if len(split) == 4 { // get queries
		dataline, result = split[1], split[2] //len(split)-2], split[:len(split)-2]
		kv := strings.Split(dataline, ":")
		o.simpleParse(strings.TrimSpace(kv[1]))
	} else { // set queries
		result = split[1]
	}

	return CommandResponse{
		Success: telnetIsResultSuccess(result),
		Raw:     r,
		Data:    o}
}

func telnetIsResultSuccess(r string) bool {
	parts := strings.Split(r, " ")
	return rawToInt(parts[1]) == 0
}

func rawToInt(raw string) int {
	stripped := strings.TrimSpace(raw)
	toint, _ := strconv.Atoi(stripped)
	return toint
}
