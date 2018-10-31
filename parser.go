package main

import (
	"fmt"
)

// Pipe ...
type Pipe interface {
	IsEmpty() bool
	In(string)
	Out() (bool, Token)
	Header() (bool, Token)
}

// TextPipe ...
type TextPipe struct {
	tokens []Token
}

// NewTextPipe ...
func NewTextPipe() *TextPipe {
	return &TextPipe{tokens: []Token{}}
}

// IsEmpty ...
func (p *TextPipe) IsEmpty() bool {
	return len(p.tokens) <= 0
}

// In ...
func (p *TextPipe) In(code string) {
	tokens := Tokenization(code)
	p.tokens = append(p.tokens, tokens...)
}

// Header ...
func (p *TextPipe) Header() (bool, Token) {
	if p.IsEmpty() {
		return false, Token{}
	}
	return true, p.tokens[0]
}

// Out ...
func (p *TextPipe) Out() (bool, Token) {
	if p.IsEmpty() {
		return false, Token{}
	}
	header := p.tokens[0]
	p.tokens = p.tokens[1:]
	return true, header
}

// Parse ...
func Parse(code string) (Expression, error) {
	p := NewTextPipe()
	p.In(code)
	ats, err := parse(p)
	if err != nil {
		return nil, fmt.Errorf("parse error %v", code)
	}
	return ats, nil
}

func parse(p Pipe) (Expression, error) {
	exp := &List{}
	ok, lp := p.Out()
	if !ok {
		return nil, fmt.Errorf("unable to find token")
	}
	if lp.Name != LPARENTHESE {
		return nil, fmt.Errorf("unable to find LPARENTHESE")
	}
	for {
		ok, token := p.Header()
		if !ok {
			return nil, fmt.Errorf("unable to find token")
		}
		switch token.Name {
		case LPARENTHESE:
			child, err := parse(p)
			if err != nil {
				return child, err
			}
			exp.PushBack(child)
		case RPARENTHESE:
			return exp, nil
		default:
			v, err := ParseToken(token)
			if err != nil {
				return v, err
			}
			exp.PushBack(v)
		}
		_, _ = p.Out()
	}
}
