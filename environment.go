package main

import (
	"fmt"
)

type frame map[Symbol]Expression

// Environment ...
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
	f := frame{
		"car":     Primitive(car),
		"cdr":     Primitive(cdr),
		"cons":    Primitive(cons),
		"list":    Primitive(list),
		"eq?":     Primitive(isEqual),
		"null?":   Primitive(isNull),
		"list?":   Primitive(isList),
		"pair?":   Primitive(isPair),
		"symbol?": Primitive(isSymbol),
		"string?": Primitive(isString),
		"p":       Primitive(printExp),
		"+":       Primitive(add),
		"-":       Primitive(subtract),
		"*":       Primitive(multiply),
		"/":       Primitive(divide),
		">":       Primitive(greaterThan),
		">=":      Primitive(greaterEqThan),
		"<":       Primitive(lessThan),
		"<=":      Primitive(lessEqThan),
		"#t":      Boolean(true),
		"#f":      Boolean(false),
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
