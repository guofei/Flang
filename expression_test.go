package main

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	if fmt.Sprintf("%v", Append(Number(1), Number(2))) != `(1 . 2)` {
		t.Error("Cons Error")
	}
	if fmt.Sprintf("%v", Append(Number(1), EmptyList())) != `(1 ())` {
		t.Error("Cons Error")
	}
	if fmt.Sprintf("%v", Append(Number(1), Append(Number(2), EmptyList()))) != `(1 (2 ()))` {
		t.Error("Cons Error")
	}
}
