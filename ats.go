package main

import (
	"container/list"
)

// Expression ...
type Expression = *list.List

// New ...
func New() Expression {
	return list.New()
}

// Variable ...
func Variable() string {
	// TODO
	return ""
}

func IsSelfEvaluating(exp Expression) bool {
	return false
}

// IsVariable ...
func IsVariable(exp Expression) bool {
	return false
}

// IsIf ...
func IsIf(exp Expression) bool {
	// TODO
	return false
}

// IsDefinition ...
func IsDefinition(exp Expression) bool {
	// TODO
	return false
}

// Operator ...
func Operator(exp Expression) Expression {
	// TODO
	return exp
}

// Operands ...
func Operands(exp Expression) Expression {
	// TODO
	return exp
}
