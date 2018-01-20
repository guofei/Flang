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
		Token{"define", FATOM},
		Token{"x", FATOM},
		Token{"(", FLPARENTHESE},
		Token{"lambda", FATOM},
		Token{"(", FLPARENTHESE},
		Token{"p1", FATOM},
		Token{")", FRPARENTHESE},
		Token{"(", FLPARENTHESE},
		Token{"y", FATOM},
		Token{"hello world", FSTRING},
		Token{"p1", FATOM},
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
