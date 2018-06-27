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

// IsAtom ...
func (p Primitive) IsAtom() bool {
	return false
}

// IsList ...
func (p Primitive) IsList() bool {
	return false
}

// String ...
func (p Primitive) String() string {
	return "primitive"
}

// Procedure ...
type Procedure struct {
	Parameters Expression
	Body       *List
	Env        *Environment
}

// IsAtom ...
func (p Procedure) IsAtom() bool {
	return false
}

// IsList ...
func (p Procedure) IsList() bool {
	return false
}

// String ...
func (p Procedure) String() string {
	return fmt.Sprintf("(lambda (%v) (%v))", p.Parameters, p.Body)
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

func firstArg(args Expression) (Expression, error) {
	list, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("args error %v", args)
	}
	return list.Car()
}

func firstArgAsList(args Expression) (*List, error) {
	list, ok := args.(*List)
	if !ok {
		return nil, fmt.Errorf("args error %v", args)
	}
	first, err := list.Car()
	if err != nil {
		return nil, fmt.Errorf("args error %v", args)
	}
	firstList, ok := first.(*List)
	if !ok {
		return nil, fmt.Errorf("args must be a list: %v", args)
	}
	return firstList, nil
}

func car(args Expression) (Expression, error) {
	list, err := firstArgAsList(args)
	if err != nil {
		return nil, err
	}
	return list.Car()
}

func cdr(args Expression) (Expression, error) {
	list, err := firstArgAsList(args)
	if err != nil {
		return nil, err
	}
	return list.Cdr()
}

func cons(args Expression) (Expression, error) {
	return args.(*List).Copy(), nil
}

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
	list, err := firstArgAsList(args)
	if err != nil {
		return nil, err
	}
	return Boolean(list.IsNull()), nil
}

func isList(args Expression) (Expression, error) {
	first, err := firstArg(args)
	if err != nil {
		return nil, err
	}
	firstList, ok := first.(*List)
	if ok {
		return Boolean(firstList.IsList()), nil
	}
	return Boolean(false), nil
}

func isPair(args Expression) (Expression, error) {
	first, err := firstArg(args)
	if err != nil {
		return nil, err
	}
	firstList, ok := first.(*List)
	if ok {
		return Boolean(firstList.IsPair()), nil
	}
	return Boolean(false), nil
}

func isSymbol(args Expression) (Expression, error) {
	first, err := firstArg(args)
	if err != nil {
		return nil, err
	}
	_, ok := first.(Symbol)
	return Boolean(ok), nil
}

func isString(args Expression) (Expression, error) {
	first, err := firstArg(args)
	if err != nil {
		return nil, err
	}
	_, ok := first.(String)
	return Boolean(ok), nil
}

func printExp(args Expression) (Expression, error) {
	first, err := firstArg(args)
	if err != nil {
		return nil, err
	}
	fmt.Println(first)
	return Symbol("OK"), nil
}
