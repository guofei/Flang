package main

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	// self evaluating
	exp := String("abc")
	res, _ := Eval(exp, nil)
	if fmt.Sprintf("%v", res) != `"abc"` {
		t.Error("Eval Error")
	}
}
