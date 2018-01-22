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
		Token{"(", LPARENTHESE},
		Token{"define", SYMBOL},
		Token{"x", SYMBOL},
		Token{"(", LPARENTHESE},
		Token{"lambda", SYMBOL},
		Token{"(", LPARENTHESE},
		Token{"p1", SYMBOL},
		Token{")", RPARENTHESE},
		Token{"(", LPARENTHESE},
		Token{"y", SYMBOL},
		Token{"\"hello world\"", STRING},
		Token{"p1", SYMBOL},
		Token{"99", NUMBER},
		Token{")", RPARENTHESE},
		Token{")", RPARENTHESE},
		Token{")", RPARENTHESE},
	}
	if !reflect.DeepEqual(tokens, except) {
		t.Error("Tokenization Error")
		t.Log("tokens: ", tokens, "except: ", except)
	}
}
