package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	code := `(+ (* 1 2) (- 2 1))`
	exp, _ := Parse(code)
	copyed := exp.(*List).Copy()
	except := `(+ ((* (1 2)) (- (2 1))))`
	if fmt.Sprintf("%v", copyed) != except {
		t.Error("Expression Error: copy")
		t.Log("exp: ", copyed, "except: ", except)
	}
}
