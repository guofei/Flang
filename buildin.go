package main

import (
	"fmt"
)

// Number ...
type Number float64

// Symbol ...
type Symbol string

// String ...
type String string

// Boolean ...
type Boolean bool

// Primitive ...
type Primitive func(Expression) (Expression, error)

// Procedure ...
type Procedure struct {
	Parameters Expression
	Body       *List
	Env        *Environment
}

// ParseToken ...
func ParseToken(token Token) (Expression, error) {
	switch token.Name {
	case NUMBER:
		return Number(ToFloat(token.Value)), nil
	case SYMBOL:
		return Symbol(token.Value), nil
	case STRING:
		return String(token.Value[1 : len(token.Value)-1]), nil
	case BOOLEAN:
		return Boolean(ToBool(token.Value)), nil
	default:
		return nil, fmt.Errorf("Can't parse token %v", token)
	}
}
