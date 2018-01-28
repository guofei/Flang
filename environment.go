package main

import (
	"fmt"
)

type frame map[Symbol]Expression

type Environment struct {
	parent *Environment
	f      frame
}

// Lookup ...
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

// Extend ...
func (base *Environment) Extend(varsExp Expression, valsExp Expression) bool {
	vars, ok := varsExp.(*List)
	if !ok {
		return false
	}
	vals, ok := valsExp.(*List)
	if !ok {
		return false
	}
	if vars.Len() != vals.Len() {
		return false
	}
	varsCar, err := vars.Car()
	if err != nil {
		return true
	}
	valsCar, err := vals.Car()
	if err != nil {
		return true
	}
	base.Set(varsCar.(Symbol), valsCar)
	varsCdr, err := vars.Cdr()
	if err != nil {
		return true
	}
	valsCdr, err := vals.Cdr()
	if err != nil {
		return true
	}
	return base.Extend(varsCdr.(*List), valsCdr.(*List))
}

// BaseEnv ...
func BaseEnv() *Environment {
	env := &Environment{nil, make(frame)}
	// TODO add more
	f := frame{
		Symbol("+"): Primitive(add),
	}
	for s, exp := range f {
		env.Set(s, exp)
	}
	return env
}

// NewChild ...
func (p *Environment) NewChild() *Environment {
	return &Environment{p, make(frame)}
}

func add(args Expression) (Expression, error) {
	argsList, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("add error %v", args)
	}
	if argsList.IsNull() {
		return Number(0), nil
	}
	car, _ := argsList.Car()
	cdr, _ := argsList.Cdr()
	n, ok := car.(Number)
	if !ok {
		return nil, fmt.Errorf("add error %v", args)
	}
	rest, err := add(cdr)
	if err != nil {
		return nil, fmt.Errorf("add error %v", args)
	}
	return n + rest.(Number), nil
}
