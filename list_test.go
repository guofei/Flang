package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	code := `(+ (* 1 2) (- 2 1))`
	exp, _ := Parse(code)
	except := `(+ ((* (1 (2 ()))) ((- (2 (1 ()))) ())))`
	if fmt.Sprintf("%v", exp.(*List).Copy()) != except {
		t.Error("Expression Error: copy")
	}
}
