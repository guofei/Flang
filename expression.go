package main

import (
	"fmt"
)

// Expression ...
type Expression interface{}

// IsSelfEvaluating ...
func IsSelfEvaluating(exp Expression) bool {
	switch exp.(type) {
	case Number, String, Boolean:
		return true
	default:
		return false
	}
}

// IsVariable ...
func IsVariable(exp Expression) bool {
	switch exp.(type) {
	case Symbol:
		return true
	default:
		return false
	}
}

// IsQuoted ...
func IsQuoted(exp Expression) bool {
	switch t := exp.(type) {
	case *Cell:
		car, err := t.Car()
		if err != nil {
			return false
		}
		switch ct := car.(type) {
		case Symbol:
			return ct == Symbol("quote")
		default:
			return false
		}
	default:
		return false
	}
}

// IsAssignment ...
func IsAssignment(exp Expression) bool {
	switch t := exp.(type) {
	case *Cell:
		car, err := t.Car()
		if err != nil {
			return false
		}
		switch ct := car.(type) {
		case Symbol:
			return ct == Symbol("set!")
		default:
			return false
		}
	default:
		return false
	}
}

// IsApplication ...
func IsApplication(exp Expression) bool {
	switch t := exp.(type) {
	case *Cell:
		return t.IsPair()
	default:
		return false
	}
}

// TextOfQuotation ...
func TextOfQuotation(exp Expression) (Expression, error) {
	c, ok := exp.(*Cell)
	if !ok {
		return nil, fmt.Errorf("unknown quatation %v", exp)
	}
	return c.Cadr()
}

// Operator ...
func Operator(exp Expression) (Expression, error) {
	c, ok := exp.(*Cell)
	if !ok {
		return nil, fmt.Errorf("unknown operator %v", exp)
	}
	return c.Car()
}

// Operands ...
func Operands(exp Expression) (Expression, error) {
	c, ok := exp.(*Cell)
	if !ok {
		return nil, fmt.Errorf("unknown operands %v", exp)
	}
	return c.Cdr()
}

// NewList ...
func EmptyList() Expression {
	return &Cell{}
}

// Cons ...
func Cons(exp1 Expression, exp2 Expression) Expression {
	res := &Cell{}
	c1, ok := exp1.(*Cell)
	if ok {
		res.car = Copy(c1)
	} else {
		res.car = exp1
	}
	c2, ok := exp2.(*Cell)
	if ok {
		res.cdr = Copy(c2)
	} else {
		res.cdr = exp2
	}
	return res
}
