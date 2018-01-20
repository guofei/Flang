package main

type Environment map[string]Expression

// LookupVariableValue ...
func LookupVariableValue(exp Expression, env Environment) (Expression, bool) {
	return exp, false
}
