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

// Out ...
func (p *IOPipe) Out() (bool, Token) {
	if p.IsEmpty() {
		fmt.Print(">*   ")
		p.s.Scan()
		p.In(p.s.Text())
	}
	header := p.tokens[0]
	p.tokens = p.tokens[1:]
	return true, header
}

// OutFirst ...
func (p *IOPipe) OutFirst() (bool, Token) {
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
		ok, lp := p.OutFirst()
		if lp.Value == "exit" {
			break
		}
		if !ok {
			continue
		}
		if lp.Name != LPARENTHESE {
			continue
		}
		ats, err := parseByPipe(p)
		if err != nil {
			continue
		}
		res, err := Eval(ats, env)
		if err != nil {
			continue
		}
		_, _ = printExp(res)
	}
}

func parseByPipe(p *IOPipe) (Expression, error) {
	exp := &List{}
loop:
	for {
		ok, token := p.Out()
		if !ok {
			return nil, fmt.Errorf("unable to find token")
		}
		switch token.Name {
		case LPARENTHESE:
			child, err := parseByPipe(p)
			if err != nil {
				return child, err
			}
			exp.PushBack(child)
		case RPARENTHESE:
			break loop
		default:
			v, err := ParseToken(token)
			if err != nil {
				return v, err
			}
			exp.PushBack(v)
		}
	}
	return exp, nil
}
