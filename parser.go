package main

// Parse ...
func Parse(code string) (Expression, error) {
	return parse(Tokenization(code))
}

func parse(tokens []Token) (Expression, error) {
	exp := &Cell{}
loop:
	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Name {
		case LPARENTHESE:
			child, err := parse(tokens[i:])
			if err != nil {
				return child, err
			}
			exp.PushBack(child)
			i += child.(*Cell).Len() + 1
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
