package main

import (
	"fmt"
	"testing"
)

func TestCons(t *testing.T) {
	exp1 := Cons(Number(1), Number(2))
	except1 := `(1 2)`
	if fmt.Sprintf("%v", exp1) != except1 {
		t.Error("Cons Error")
		t.Log("exp: ", exp1, "except: ", except1)
	}
	exp2 := Cons(Number(1), EmptyList())
	except2 := `1`
	if fmt.Sprintf("%v", exp2) != `1` {
		t.Error("Cons Error")
		t.Log("exp: ", exp2, "except: ", except2)
	}
	exp3 := Cons(Number(1), Cons(Number(2), EmptyList()))
	except3 := `(1 2)`
	if fmt.Sprintf("%v", exp3) != except3 {
		t.Error("Cons Error")
		t.Log("exp: ", exp3, "except: ", except3)
	}
}
