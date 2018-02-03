package main

// Pipe ...
type Pipe struct {
	tokens []Token
}

// NewPipe ...
func NewPipe() *Pipe {
	return &Pipe{tokens: []Token{}}
}

// IsEmpty ...
func (p *Pipe) IsEmpty() bool {
	return len(p.tokens) <= 0
}

// In ...
func (p *Pipe) In(code string) {
	tokens := Tokenization(code)
	p.tokens = append(p.tokens, tokens...)
}

// Out ...
func (p *Pipe) Out() (bool, Token) {
	if p.IsEmpty() {
		return false, Token{}
	}
	header := p.tokens[0]
	p.tokens = p.tokens[1:]
	return true, header
}
