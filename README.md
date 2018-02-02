# Flang

A Scheme dialect written in Golang

## Concept

[The eval-apply cycle](http://sarabander.github.io/sicp/html/4_002e1.xhtml#g_t4_002e1_002e4)

## Hello World

``` lisp
(p "hello world")
```

## Basic types

Number, Symbol, String, Boolean, Lambda, List, Pair

## Definition

### Define a varable

``` lisp
(define <var> <value>)
```

Example

``` lisp
(define x 1)
```

### Define a function

``` lisp
(define (<var> <param1> ... <paramN>) <value>)
```

Example

``` lisp

(define (factorial n)
  (if (eq? n 1)
      1
      (* (factorial (- n 1)) n)))
(factorial 5)

;; result: 120
```

## Assignment

``` lisp
(set! <var> <value>)
```

Example

``` lisp
(define x 1)
(set! x 10)
(p x)

;; result: 10
```

## Conditions

if, cond

Example

``` lisp
(define x 1)
(if (> x 5) 
    (p "greater than 5")
    (p "less than 5"))
```

## Function

### Lambda
``` lisp
((lambda (x) (* x x)) 2)
```

### Standard Library

cons, car, cdr, list, pair ...

## TODOS

- [ ] REPL
- [ ] Macro
- [ ] Comment
- [ ] Bootstrapping

