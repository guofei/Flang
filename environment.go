package main

import (
	"fmt"
)

type frame map[Symbol]Expression

type Environment struct {
	parent *Environment
	f      frame
}

// LookupVariableValue ...
func LookupVariableValue(exp Expression, env *Environment) (Expression, error) {
	s, ok := exp.(Symbol)
	if !ok {
		return nil, fmt.Errorf("unbound variable %v", exp)
	}
	v, ok := env.f[s]
	if ok {
		return v, nil
	}
	if env.parent == nil {
		return nil, fmt.Errorf("unbound variable %v", exp)
	}
	return LookupVariableValue(exp, env.parent)
}

// EvalAssignment ...
func EvalAssignment(exp Expression, env *Environment) (Expression, error) {
	c, ok := exp.(*Cell)
	if !ok {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	variable, err := c.Cadr()
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	valueBody, err := c.Caddr()
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	value, err := Eval(valueBody, env)
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	v, ok := variable.(Symbol)
	if !ok {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	env.f[v] = value
	return value, nil
}
