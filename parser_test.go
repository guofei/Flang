package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	code := `(+ (* 1 1) (- 3 1))`
	l := Parse(code)
	fmt.Println(l)
}
