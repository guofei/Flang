package main

import (
	"fmt"
)

// Apply ...
func Apply(procedure Expression, args Expression) (Expression, error) {
	switch p := procedure.(type) {
	case Primitive:
		return p(args)
		/*
			case Procedure:
				// TODO
				return nil, nil
		*/
	default:
		return nil, fmt.Errorf("unknown procedure type: Apply %v", procedure)
	}
}
