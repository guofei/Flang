package main

func Parse(code string) Expression {
	tokens := Tokenization(code)
	return parse(tokens)
}

func parse(tokens []Token) Expression {
	exp := &Cell{}
loop:
	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Name {
		case FLPARENTHESE:
			child := parse(tokens[i:])
			exp.PushBack(child)
			i += child.(*Cell).Len() + 1
		case FRPARENTHESE:
			break loop
		default:
			exp.PushBack(ParseToken(token))
		}
	}

	return exp
}
