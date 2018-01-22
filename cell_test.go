package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	code := `(+ (* 1 2) (- 2 1))`
	exp, _ := Parse(code)
	except := `(+ ((* (1 (2 ()))) ((- (2 (1 ()))) ())))`
	if fmt.Sprintf("%v", Copy(exp.(*Cell))) != except {
		t.Error("Expression Error: copy")
	}
}
