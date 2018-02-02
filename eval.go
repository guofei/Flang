package main

import (
	"fmt"
)

// Eval ...
func Eval(exp Expression, env *Environment) (Expression, error) {
	switch {
	case isSelfEvaluating(exp):
		return exp, nil
	case isVariable(exp):
		return env.Lookup(exp.(Symbol))
	case isSpecialForm(exp):
		return evalSpecialForm(exp.(*List), env)
	case isApplication(exp):
		list := exp.(*List)
		operator, _ := list.Car()
		procedure, err := Eval(operator, env)
		if err != nil {
			return nil, fmt.Errorf("unknown operator %v", operator)
		}
		operands, _ := list.Cdr()
		args, err := listOfValues(operands, env)
		if err != nil {
			return nil, fmt.Errorf("unknown operands %v", operands)
		}
		return Apply(procedure, args)
	default:
		return nil, fmt.Errorf("unknown expression type %v", exp)
	}
}

func isSelfEvaluating(exp Expression) bool {
	switch exp.(type) {
	case Number, String, Boolean:
		return true
	default:
		return false
	}
}

func isVariable(exp Expression) bool {
	switch exp.(type) {
	case Symbol:
		return true
	default:
		return false
	}
}

func isApplication(exp Expression) bool {
	switch t := exp.(type) {
	case *List:
		return t.IsPair()
	default:
		return false
	}
}

func isSpecialForm(exp Expression) bool {
	switch t := exp.(type) {
	case *List:
		car, err := t.Car()
		if err != nil {
			return false
		}
		switch ct := car.(type) {
		case Symbol:
			_, ok := formFunc(ct)
			return ok
		default:
			return false
		}
	default:
		return false
	}
}

func listOfValues(exps Expression, env *Environment) (Expression, error) {
	if exps.(*List).IsNull() {
		return EmptyList(), nil
	}
	car, _ := exps.(*List).Car()
	firstValue, err := Eval(car, env)
	if err != nil {
		return firstValue, err
	}
	cdr, _ := exps.(*List).Cdr()
	restValues, err := listOfValues(cdr, env)
	if err != nil {
		return restValues, err
	}
	return Append(firstValue, restValues), nil
}

func formFunc(f Symbol) (func(*List, *Environment) (Expression, error), bool) {
	dict := map[Symbol]func(*List, *Environment) (Expression, error){
		"quote":  textOfQuotation,
		"set!":   evalAssignment,
		"define": evalDefinition,
		"if":     evalIf,
		"lambda": evalLambda,
		"begin":  evalBegin,
		"cond":   evalCond,
	}
	v, ok := dict[f]
	return v, ok
}

func evalSpecialForm(exp *List, env *Environment) (Expression, error) {
	form, _ := exp.Car()
	fun, _ := formFunc(form.(Symbol))
	return fun(exp, env)
}

func textOfQuotation(exp *List, env *Environment) (Expression, error) {
	return exp.Cadr()
}

func evalAssignment(exp *List, env *Environment) (Expression, error) {
	variable, err := exp.Cadr()
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	valueBody, err := exp.Caddr()
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	value, err := Eval(valueBody, env)
	if err != nil {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	k, ok := variable.(Symbol)
	if !ok {
		return nil, fmt.Errorf("assignment error %v", exp)
	}
	env.Set(k, value)
	return Symbol("ok"), nil
}

func definitionVariable(exp *List) (Symbol, error) {
	variableExp, err := exp.Cadr()
	if err != nil {
		return "", fmt.Errorf("definition error %v", exp)
	}
	switch t := variableExp.(type) {
	case Symbol:
		return t, nil
	default:
		v, err := variableExp.(*List).Car()
		if err != nil {
			return "", fmt.Errorf("definition error %v", variableExp)
		}
		s, ok := v.(Symbol)
		if !ok {
			return "", fmt.Errorf("definition error %v", variableExp)
		}
		return s, nil
	}
}

func makeLambda(params Expression, body Expression) *List {
	return Append(Symbol("lambda"), Append(params, body))
}

func definitionValue(exp *List) (Expression, error) {
	variableExp, err := exp.Cadr()
	if err != nil {
		return nil, fmt.Errorf("definition error %v", exp)
	}
	switch t := variableExp.(type) {
	case Symbol:
		value, err := exp.Caddr()
		if err != nil {
			return nil, fmt.Errorf("definition error %v", exp)
		}
		return value, err
	default:
		params, err := t.(*List).Cdr()
		if err != nil {
			return nil, fmt.Errorf("definition error %v", exp)
		}
		body, err := exp.Cddr()
		if err != nil {
			return nil, fmt.Errorf("definition error %v", exp)
		}
		return makeLambda(params, body), nil
	}
}

func evalDefinition(exp *List, env *Environment) (Expression, error) {
	variable, err := definitionVariable(exp)
	if err != nil {
		return nil, fmt.Errorf("definition error %v", exp)
	}
	value, err := definitionValue(exp)
	if err != nil {
		return nil, fmt.Errorf("definition error %v", exp)
	}
	v, err := Eval(value, env)
	if err != nil {
		return nil, fmt.Errorf("definition error %v", exp)
	}
	env.Set(variable, v)
	return Symbol("ok"), nil
}

func ifAlternative(exp *List) (Expression, error) {
	cdddr, err := exp.Cdddr()
	if err != nil {
		return nil, fmt.Errorf("if error %v", exp)
	}
	switch t := cdddr.(type) {
	case *List:
		if t.IsNull() {
			return Boolean(false), nil
		}
		return t.Car()
	default:
		return nil, fmt.Errorf("if error %v", exp)
	}
}

func evalIf(exp *List, env *Environment) (Expression, error) {
	predicateExp, err := exp.Cadr()
	if err != nil {
		return nil, fmt.Errorf("if error %v", exp)
	}
	predicate, err := Eval(predicateExp, env)
	if err != nil {
		return nil, fmt.Errorf("if error %v", exp)
	}
	if IsTrue(predicate) {
		consequent, e := exp.Caddr()
		if e != nil {
			return nil, fmt.Errorf("if error %v", exp)
		}
		return Eval(consequent, env)
	}
	alternative, err := ifAlternative(exp)
	if err != nil {
		return nil, fmt.Errorf("if error %v", exp)
	}
	return Eval(alternative, env)
}

func evalLambda(exp *List, env *Environment) (Expression, error) {
	params, err := exp.Cadr()
	if err != nil {
		return nil, fmt.Errorf("lambda error %v", exp)
	}
	body, err := exp.Cddr()
	if err != nil {
		return nil, fmt.Errorf("lambda error %v", exp)
	}
	return Procedure{params, body.(*List), env}, nil
}

func evalBegin(exps *List, env *Environment) (Expression, error) {
	cdr, err := exps.Cdr()
	if err != nil {
		return nil, fmt.Errorf("begin error %v", exps)
	}
	return EvalSequence(cdr.(*List), env)
}

// EvalSequence ...
func EvalSequence(exps *List, env *Environment) (Expression, error) {
	car, err := exps.Car()
	if err != nil {
		return nil, fmt.Errorf("sequence error 1 %v", exps)
	}
	cdr, err := exps.Cdr()
	if err != nil {
		return nil, fmt.Errorf("sequence error 2 %v", exps)
	}
	if cdr.(*List).IsNull() {
		return Eval(car, env)
	}
	_, err = Eval(car, env)
	if err != nil {
		return nil, fmt.Errorf("sequence error 3 %v", car)
	}
	return EvalSequence(cdr.(*List), env)
}

// TODO
func evalCond(exp *List, env *Environment) (Expression, error) {
	return nil, nil
}
