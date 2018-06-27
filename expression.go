package main

import (
	"fmt"
)

// Expression ...
type Expression interface {
	IsAtom() bool
	IsList() bool
	String() string
}

// IsTrue ...
func IsTrue(exp Expression) bool {
	switch t := exp.(type) {
	case Boolean:
		return bool(t)
	default:
		return true
	}
}

// Cons ...
func Cons(args ...Expression) *List {
	res := &List{}
	c1, ok := args[0].(*List)
	if ok {
		res.car = c1
	} else {
		res.car = args[0]
	}
	c2, ok := args[1].(*List)
	if ok {
		res.cdr = c2
	} else {
		res.cdr = &List{args[1], &List{}}
	}
	return res
}
