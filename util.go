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

// IsBool ...
func IsBool(s string) bool {
	return s == "#t" || s == "#f"
}

// ToFloat ...
func ToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// ToBool ...
func ToBool(s string) bool {
	return s == "#t"
}
