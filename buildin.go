package main

type Expression interface{}

type Number float64

type Symbol string

type String string

type Boolean bool

// ParseToken ...
func ParseToken(token Token) Expression {
	switch token.Name {
	case FNUMBER:
		return Number(ToFloat(token.Value))
	case FSYMBOL:
		return Symbol(token.Value)
	case FSTRING:
		return String(token.Value[1 : len(token.Value)-1])
	case FBOOLEAN:
		return Boolean(ToBool(token.Value))
	default:
		return nil
	}
}
