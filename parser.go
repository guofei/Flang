package main

func Parse(code string) Expression {
	tokens := Tokenization(code)
	return parse(tokens)
}

// TODO
func parse(tokens []Token) Expression {
	exp := New()
loop:
	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case FLPARENTHESE:
			child := parse(tokens[i:])
			exp.PushBack(child)
			i += child.Len() + 1
		case FRPARENTHESE:
			break loop
		default:
			exp.PushBack(token)
		}
	}

	return exp
}
