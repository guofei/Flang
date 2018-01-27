package main

import (
	"fmt"
)

type frame map[Symbol]Expression

type Environment struct {
	parent *Environment
	f      frame
}

// Lookup
func (env *Environment) Lookup(s Symbol) (Expression, error) {
	v, ok := env.f[s]
	if ok {
		return v, nil
	}
	if env.parent == nil {
		return nil, fmt.Errorf("unbound variable %v", s)
	}
	return env.parent.Lookup(s)
}

// Set ...
func (env *Environment) Set(k Symbol, v Expression) {
	env.f[k] = v
}
