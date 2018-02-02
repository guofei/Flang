package main

import (
	"fmt"
	"testing"
)

func TestEvalSelfEval(t *testing.T) {
	exp := String("abc")
	res, _ := Eval(exp, nil)
	if fmt.Sprintf("%v", res) != `"abc"` {
		t.Error("Eval Error")
	}
}

func TestEvalPrimitiveAdd(t *testing.T) {
	env := BaseEnv()
	code := `(+ 1 2)`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(3) {
		t.Error("Eval Error")
	}
}

func TestEvalPrimitiveMultiply(t *testing.T) {
	env := BaseEnv()
	code := `(* 2 2)`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(4) {
		t.Error("Eval Error")
	}
}

func TestEvalDefinition(t *testing.T) {
	env := BaseEnv()
	code := `(define x 1)`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Symbol("ok") {
		t.Error("Eval Error")
	}
}

func TestEvalDefinitionfunction(t *testing.T) {
	env := BaseEnv()
	code := `(define (double x) (+ x x))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Symbol("ok") {
		t.Error("Eval Error")
	}
}

func TestEvalBegin(t *testing.T) {
	env := BaseEnv()
	code := `(begin
                   (define a 3)
                   (define b 5)
                   (+ a b))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(8) {
		t.Error("Eval Error")
	}
}

func TestEvalApplication(t *testing.T) {
	env := BaseEnv()
	code := `(begin
                   (define (double x) (+ x x))
                   (define b 5)
                   (+ (double 2) b))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(9) {
		t.Error("Eval Error")
	}
}

func TestEvalIf(t *testing.T) {
	env := BaseEnv()
	code := `(if (eq? 1 1)
                     (+ 1 1)
                     (- 1 1))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(2) {
		t.Error("Eval Error")
	}
}

func TestEvalElse(t *testing.T) {
	env := BaseEnv()
	code := `(if (eq? 2 1)
                     (+ 1 1)
                     (- 1 1))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(0) {
		t.Error("Eval Error")
	}
}

func TestEvalIfAndDefine(t *testing.T) {
	env := BaseEnv()
	code := `(begin
                   (define (f n)
                     (if (eq? n 1)
                         1
                         (* 2 n)))
                   (f 5))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(10) {
		t.Error("Eval Error")
	}
}

func TestEvalRecursion(t *testing.T) {
	env := BaseEnv()
	code := `(begin
                   (define (factorial n)
                     (if (eq? n 1)
                         1
                         (* (factorial (- n 1)) n)))
                   (factorial 5))`
	ats, _ := Parse(code)
	res, _ := Eval(ats, env)
	if res != Number(120) {
		t.Error("Eval Error")
	}
}
