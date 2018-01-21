package main

type TokenName int

type Token struct {
	Value string
	Name  TokenName
}

const (
	FLPARENTHESE TokenName = iota
	FRPARENTHESE
	FSYMBOL
	FBOOLEAN
	FNUMBER
	FSTRING
)

func Tokenization(code string) []Token {
	tokens := []Token{}
	i := 0
	for {
		i = trimLeft(code, i)
		if i < 0 {
			break
		}
		item, nextIndex := nextItem(code, i)
		if nextIndex < 0 {
			break
		}
		tokens = append(tokens, Token{item, getType(item)})
		i = nextIndex
	}
	return tokens
}

func isSpliter(r byte) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

func getString(code string, i int) (string, int) {
	if i+1 >= len(code) {
		return "", -1
	}
	item := code[i : i+1]
	next := i + 1
	for {
		if next >= len(code) {
			return "", -1
		}
		item += code[next : next+1]
		if code[next] == '"' {
			next++
			break
		}
		next++
	}
	return item, next
}

func trimLeft(code string, i int) int {
	if i >= len(code) {
		return -1
	}
	next := i
	for isSpliter(code[next]) {
		next++
		if next >= len(code) {
			return -1
		}
	}
	return next
}

func nextItem(code string, i int) (string, int) {
	if i >= len(code) {
		return "", -1
	}
	if code[i] == '(' || code[i] == ')' {
		return code[i : i+1], i + 1
	}
	if code[i] == '"' {
		return getString(code, i)
	}
	var item string
	next := i
	for {
		if isSpliter(code[next]) || code[next] == '(' || code[next] == ')' {
			break
		}
		item += code[next : next+1]
		next++
		if next >= len(code) {
			return "", -1
		}
	}
	return item, next
}

func getType(item string) TokenName {
	switch {
	case item == "(":
		return FLPARENTHESE
	case item == ")":
		return FRPARENTHESE
	case IsBool(item):
		return FBOOLEAN
	case IsString(item):
		return FSTRING
	case IsNumeric(item):
		return FNUMBER
	default:
		return FSYMBOL
	}
}
