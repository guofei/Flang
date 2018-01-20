package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	code := `(+ (* 1 1) (- 3 1))`
	l := Parse(code)
	for e := l.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case Expression:
			l2 := e.Value.(Expression)
			for e2 := l2.Front(); e2 != nil; e2 = e2.Next() {
				fmt.Println(e2.Value)
			}
		default:
			fmt.Println(e.Value)
		}
	}
}
