package main

import (
	"bufio"
	"fmt"
	"os"
)

// IOPipe ...
type IOPipe struct {
	tokens []Token
	s      *bufio.Scanner
}

// NewIOPipe ...
func NewIOPipe() *IOPipe {
	return &IOPipe{tokens: []Token{}, s: bufio.NewScanner(os.Stdin)}
}

// IsEmpty ...
func (p *IOPipe) IsEmpty() bool {
	return len(p.tokens) <= 0
}

// In ...
func (p *IOPipe) In(code string) {
	tokens := Tokenization(code)
	p.tokens = append(p.tokens, tokens...)
}

// Header ...
func (p *IOPipe) Header() (bool, Token) {
	if p.IsEmpty() {
		fmt.Print(">> ")
		p.s.Scan()
		p.In(p.s.Text())
	}
	return true, p.tokens[0]
}

// Out ...
func (p *IOPipe) Out() (bool, Token) {
	if p.IsEmpty() {
		fmt.Print(">> ")
		p.s.Scan()
		p.In(p.s.Text())
	}
	header := p.tokens[0]
	p.tokens = p.tokens[1:]
	return true, header
}

// REPL ...
func REPL() {
	p := NewIOPipe()
	fmt.Println("Welcome to Flang")
	env := BaseEnv()
	for {
		ats, err := parse(p)
		if err != nil {
			continue
		}
		res, err := Eval(ats, env)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, _ = printExp(res)
	}
}
