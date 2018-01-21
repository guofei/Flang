package main

import (
	"reflect"
	"testing"
)

func TestTokenization(t *testing.T) {
	code := `
                 (define x
                   ( lambda (p1)
                     (y "hello world" p1 99)))
                `
	tokens := Tokenization(code)
	except := []Token{
		Token{"(", FLPARENTHESE},
		Token{"define", FSYMBOL},
		Token{"x", FSYMBOL},
		Token{"(", FLPARENTHESE},
		Token{"lambda", FSYMBOL},
		Token{"(", FLPARENTHESE},
		Token{"p1", FSYMBOL},
		Token{")", FRPARENTHESE},
		Token{"(", FLPARENTHESE},
		Token{"y", FSYMBOL},
		Token{"\"hello world\"", FSTRING},
		Token{"p1", FSYMBOL},
		Token{"99", FNUMBER},
		Token{")", FRPARENTHESE},
		Token{")", FRPARENTHESE},
		Token{")", FRPARENTHESE},
	}
	if !reflect.DeepEqual(tokens, except) {
		t.Error("Tokenization Error")
		t.Log("tokens: ", tokens, "except: ", except)
	}
}
