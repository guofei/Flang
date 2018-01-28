package main

import (
	"fmt"
)

// Apply ...
func Apply(procedure Expression, args Expression) (Expression, error) {
	switch p := procedure.(type) {
	case Primitive:
		return p(args)
	case Procedure:
		childEnv := p.Env.NewChild()
		childEnv.Extend(p.Parameters, args)
		return EvalSequence(p.Body, childEnv)
	default:
		return nil, fmt.Errorf("unknown procedure type %v", procedure)
	}
}
