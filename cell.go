package main

import (
	"fmt"
)

// TODO: Change rescursion to for loop

// Cell ...
type Cell struct {
	Car Expression
	Cdr Expression
}

// String ...
func (cell *Cell) String() string {
	switch {
	case cell.IsNull():
		return ""
	case cell.Len() == 1:
		return fmt.Sprintf("%v", cell.Car)
	case cell.IsList():
		return fmt.Sprintf("(%v %v)", cell.Car, cell.Cdr)
	default:
		return fmt.Sprintf("(%v . %v)", cell.Car, cell.Cdr)
	}
}

// IsNull ...
func (cell *Cell) IsNull() bool {
	return cell.Car == nil && cell.Cdr == nil
}

// IsList ...
func (cell *Cell) IsList() bool {
	if cell.IsNull() {
		return true
	}
	switch cell.Cdr.(type) {
	case *Cell:
		return cell.Cdr.(*Cell).IsList()
	default:
		return false
	}
}

// Len ...
func (cell *Cell) Len() int {
	if cell.IsNull() {
		return 0
	}
	return 1 + cell.Cdr.(*Cell).Len()
}

// PushBack ...
func (cell *Cell) PushBack(exp Expression) {
	if cell.IsNull() {
		cell.Car = exp
		cell.Cdr = &Cell{}
	} else {
		cell.Cdr.(*Cell).PushBack(exp)
	}
}
