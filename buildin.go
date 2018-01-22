package main

import (
	"fmt"
)

type Number float64

type Symbol string

type String string

type Boolean bool

// String ...
func (s String) String() string {
	return fmt.Sprintf("\"%v\"", string(s))
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
