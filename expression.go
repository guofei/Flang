package main

import (
	"fmt"
)

// Expression ...
type Expression interface {
	String() string
}

// String ...
func (n Number) String() string {
	return fmt.Sprintf("%v", float64(n))
}

// String ...
func (s Symbol) String() string {
	return string(s)
}

// String ...
func (s String) String() string {
	return fmt.Sprintf("\"%v\"", string(s))
}

// String ...
func (b Boolean) String() string {
	if b {
		return "#t"
	}
	return "#f"
}

// String ...
func (f Primitive) String() string {
	return "primitive"
}

// String ...
func (list *List) String() string {
	if list.Len() <= 0 {
		return "()"
	}
	return fmt.Sprintf("(%v)", list.childToString())
}

// childToString ...
func (list *List) childToString() string {
	if list.Len() == 0 {
		return ""
	}
	if list.Len() == 1 {
		return fmt.Sprintf("%v", list.car)
	}

	switch t := list.cdr.(type) {
	case *List:
		return fmt.Sprintf("%v %v", list.car, t.childToString())
	default:
		return fmt.Sprintf("%v %v", list.car, t)
	}
}

// String ...
func (p Procedure) String() string {
	return fmt.Sprintf("(lambda (%v) (%v))", p.Parameters, p.Body)
}

// EmptyList ...
func EmptyList() *List {
	return &List{}
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
