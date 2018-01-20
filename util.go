package main

import (
	"strconv"
)

// IsNumeric ...
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsString ...
func IsString(s string) bool {
	return s[0] == '"' && s[len(s)-1] == '"'
}
