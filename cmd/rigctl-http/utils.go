package main

import (
  "strconv"
  "strings"
)

// check if str exists in list []string
// often used to test against options available on the rig
func hasOption(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// convert string to int
// telnet always returns strings but the api deals in ints where appropriate
func strToInt(raw string) int {
	stripped := strings.TrimSpace(raw)
	toint, _ := strconv.Atoi(stripped)
	return toint
}
