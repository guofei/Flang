package main

import (
	"fmt"
	"testing"
)

func TestBuildInCons(t *testing.T) {
	env := BaseEnv()
	code := `(cons 1 (list 2 3))`
	exp, _ := Parse(code)
	res, _ := Eval(exp, env)
	except := `(1 (2 3))`
	if fmt.Sprintf("%v", res) != except {
		t.Error("Expression Error: buildin cons")
		t.Log("cons: ", res, "except: ", except)
	}
}

func TestBuildInList(t *testing.T) {
	env := BaseEnv()
	code := `(list 1 2 3)`
	exp, _ := Parse(code)
	res, _ := Eval(exp, env)
	except := `(1 2 3)`
	if fmt.Sprintf("%v", res) != except {
		t.Error("Expression Error: buildin list")
		t.Log("list: ", res, "except: ", except)
	}
}
