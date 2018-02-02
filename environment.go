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

func subtract(args Expression) (Expression, error) {
	argsList, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("subtract error %v", args)
	}
	if argsList.IsNull() {
		return nil, fmt.Errorf("subtract error %v", args)
	}
	car, _ := argsList.Car()
	n, ok := car.(Number)
	if !ok {
		return nil, fmt.Errorf("subtract error %v", args)
	}
	cdr, _ := argsList.Cdr()
	sum, err := add(cdr)
	if err != nil {
		return nil, fmt.Errorf("subtract error %v", args)
	}
	return n - sum.(Number), nil
}

func multiply(args Expression) (Expression, error) {
	argsList, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("multiply error %v", args)
	}
	if argsList.IsNull() {
		return Number(1), nil
	}
	car, _ := argsList.Car()
	cdr, _ := argsList.Cdr()
	n, ok := car.(Number)
	if !ok {
		return nil, fmt.Errorf("multiply error %v", args)
	}
	rest, err := multiply(cdr)
	if err != nil {
		return nil, fmt.Errorf("multiply error %v", args)
	}
	return n * rest.(Number), nil
}

func comparable(args Expression) bool {
	list, ok := args.(*List)
	if !ok {
		return false
	}
	car, err := list.Car()
	if err != nil {
		return false
	}
	cadr, err := list.Cadr()
	if err != nil {
		return false
	}
	_, ok = car.(Number)
	if !ok {
		return false
	}
	_, ok = cadr.(Number)
	return ok
}

func divide(args Expression) (Expression, error) {
	ok := comparable(args)
	if !ok {
		return nil, fmt.Errorf("divide error %v", args)
	}
	list, _ := args.(*List)
	n1, _ := list.Car()
	n2, _ := list.Cadr()
	if !ok {
		return nil, fmt.Errorf("divide error %v", args)
	}
	return n1.(Number) / n2.(Number), nil
}

func greaterThan(args Expression) (Expression, error) {
	ok := comparable(args)
	if !ok {
		return nil, fmt.Errorf("> error %v", args)
	}
	list, _ := args.(*List)
	n1, _ := list.Car()
	n2, _ := list.Cadr()
	if !ok {
		return nil, fmt.Errorf("> error %v", args)
	}
	return Boolean(n1.(Number) > n2.(Number)), nil
}

func lessThan(args Expression) (Expression, error) {
	ok := comparable(args)
	if !ok {
		return nil, fmt.Errorf("< error %v", args)
	}
	list, _ := args.(*List)
	n1, _ := list.Car()
	n2, _ := list.Cadr()
	if !ok {
		return nil, fmt.Errorf("< error %v", args)
	}
	return Boolean(n1.(Number) < n2.(Number)), nil
}

func greaterEqThan(args Expression) (Expression, error) {
	ok := comparable(args)
	if !ok {
		return nil, fmt.Errorf(">= error %v", args)
	}
	list, _ := args.(*List)
	n1, _ := list.Car()
	n2, _ := list.Cadr()
	if !ok {
		return nil, fmt.Errorf(">= error %v", args)
	}
	return Boolean(n1.(Number) >= n2.(Number)), nil
}

func lessEqThan(args Expression) (Expression, error) {
	ok := comparable(args)
	if !ok {
		return nil, fmt.Errorf("<= error %v", args)
	}
	list, _ := args.(*List)
	n1, _ := list.Car()
	n2, _ := list.Cadr()
	if !ok {
		return nil, fmt.Errorf("<= error %v", args)
	}
	return Boolean(n1.(Number) <= n2.(Number)), nil
}

func car(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("car error %v", args)
	}
	return list.Car()
}

func cdr(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("cdr error %v", args)
	}
	return list.Cdr()
}

// TODO test
func cons(args Expression) (Expression, error) {
	return args.(*List).Copy(), nil
}

// TODO test
func list(args Expression) (Expression, error) {
	return args.(*List).Copy(), nil
}

func isEqual(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return Boolean(false), nil
	}
	car, err := list.Car()
	if err != nil {
		return nil, err
	}
	cadr, err := list.Cadr()
	if err != nil {
		return nil, err
	}
	return Boolean(car == cadr), nil
}

func isNull(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return Boolean(false), nil
	}
	return Boolean(list.IsNull()), nil
}

func isList(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return Boolean(false), nil
	}
	return Boolean(list.IsList()), nil
}

func isPair(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return Boolean(false), nil
	}
	return Boolean(list.IsPair()), nil
}

func isSymbol(args Expression) (Expression, error) {
	_, ok := args.(Symbol)
	return Boolean(ok), nil
}

func isString(args Expression) (Expression, error) {
	_, ok := args.(String)
	return Boolean(ok), nil
}

func printExp(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("print error %v", args)
	}
	car, _ := list.Car()
	fmt.Println(car)
	return Symbol("OK"), nil
}
