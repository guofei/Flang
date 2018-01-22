package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	code := `(+ (* 1 1) (f "xx" "yy"))`
	ats, _ := Parse(code)
	except := `(+ ((* (1 (1 ()))) ((f ("xx" ("yy" ()))) ())))`
	if fmt.Sprintf("%v", ats) != except {
		t.Error("Parse Error")
		t.Log("ats: ", ats, "except: ", except)
	}
}
