package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	code := `(+ (* 1 1) (f "xx" "yy"))`
	ats, _ := Parse(code)
	except := `(+ ((* (1 1)) (f ("xx" "yy"))))`
	if fmt.Sprintf("%v", ats) != except {
		t.Error("Parse Error")
		t.Log("ats: ", ats, "except: ", except)
	}
}

func TestParseDefinition(t *testing.T) {
	code := `(begin
                   (define (double x) (+ x x))
                   (define b 5)
                   (+ (double 2) b))`
	ats, _ := Parse(code)
	except := `(begin ((define ((double x) (+ (x x)))) ((define (b 5)) (+ ((double 2) b)))))`
	if fmt.Sprintf("%v", ats) != except {
		t.Error("Parse Error")
		t.Log("ats: ", ats, "except: ", except)
	}
}
