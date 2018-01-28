package main

import (
	"fmt"
)

// Parse ...
func Parse(code string) (Expression, error) {
	exp, i := parse(Tokenization(code))
	if i < 0 {
		return nil, fmt.Errorf("parse error %v", code)
	}
	return exp, nil
}

func parse(tokens []Token) (Expression, int) {
	exp := &List{}
	i := 1
loop:
	for ; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Name {
		case LPARENTHESE:
			child, count := parse(tokens[i:])
			if count < 0 {
				return child, count
			}
			exp.PushBack(child)
			i += count
		case RPARENTHESE:
			break loop
		default:
			v, err := ParseToken(token)
			if err != nil {
				return v, -1
			}
			exp.PushBack(v)
		}
	}

	return exp, i
}
