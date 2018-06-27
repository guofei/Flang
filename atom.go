package main

import (
	"fmt"
)

// String ...
func (n Number) String() string {
	return fmt.Sprintf("%v", float64(n))
}

// String ...
func (s Symbol) String() string {
	return string(s)
}

// String ...
func (s String) String() string {
	return fmt.Sprintf("\"%v\"", string(s))
}

// String ...
func (b Boolean) String() string {
	if b {
		return "#t"
	}
	return "#f"
}

// IsAtom ...
func (n Number) IsAtom() bool {
	return true
}

// IsAtom ...
func (s Symbol) IsAtom() bool {
	return true
}

// IsAtom ...
func (s String) IsAtom() bool {
	return true
}

// IsAtom ...
func (b Boolean) IsAtom() bool {
	return true
}

// IsList ...
func (n Number) IsList() bool {
	return false
}

// IsList ...
func (s Symbol) IsList() bool {
	return false
}

// IsList ...
func (s String) IsList() bool {
	return false
}

// IsList ...
func (b Boolean) IsList() bool {
	return false
}
