package main

import (
	"fmt"
	"testing"
)

func TestIsQuoted(t *testing.T) {
	code := `(quote 1 2 3)`
	exp, _ := Parse(code)
	if !IsQuoted(exp) {
		t.Error("Expression Error: quote")
	}
}

func TestCons(t *testing.T) {
	if fmt.Sprintf("%v", Cons(1, 2)) != `(1 . 2)` {
		t.Error("Cons Error")
	}
	if fmt.Sprintf("%v", Cons(1, EmptyList())) != `(1 ())` {
		t.Error("Cons Error")
	}
	if fmt.Sprintf("%v", Cons(1, Cons(2, EmptyList()))) != `(1 (2 ()))` {
		t.Error("Cons Error")
	}
}
