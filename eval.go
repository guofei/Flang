package main

import (
	"fmt"
)

// Eval ...
func Eval(exp Expression, env *Environment) (Expression, error) {
	// TODO
	switch {
	case IsSelfEvaluating(exp):
		return exp, nil
	case IsVariable(exp):
		return LookupVariableValue(exp, env)
	case IsQuoted(exp):
		return TextOfQuotation(exp)
	case IsApplication(exp):
		// return Apply(Eval(exp.(*Cell).Car(), env))
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown expression type %v", exp)
	}
}

// ListOfValues ...
func ListOfValues(exps Expression, env *Environment) (Expression, error) {
	if exps.(*Cell).IsNull() {
		return EmptyList(), nil
	}
	firstValue, err := Eval(exps, env)
	if err != nil {
		return firstValue, err
	}
	return nil, nil
}
