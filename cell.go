package main

import (
	"fmt"
)

// TODO: Change rescursion to for loop

// Cell ...
type Cell struct {
	car Expression
	cdr Expression
}

// Car ...
func (cell *Cell) Car() (Expression, error) {
	if cell.IsNull() {
		return nil, fmt.Errorf("attempt to apply car on nil %v", cell)
	}
	return cell.car, nil
}

// Cdr ...
func (cell *Cell) Cdr() (Expression, error) {
	if cell.IsNull() {
		return nil, fmt.Errorf("attempt to apply cdr on nil %v", cell)
	}
	return cell.cdr, nil
}

// Cadr second
func (cell *Cell) Cadr() (Expression, error) {
	cdr, err := cell.Cdr()
	if err != nil {
		return nil, err
	}
	return cdr.(*Cell).Car()
}

// Cddr ...
func (cell *Cell) Cddr() (Expression, error) {
	cdr, err := cell.Cdr()
	if err != nil {
		return nil, err
	}
	return cdr.(*Cell).Cdr()
}

// Caddr third
func (cell *Cell) Caddr() (Expression, error) {
	cddr, err := cell.Cddr()
	if err != nil {
		return nil, err
	}
	return cddr.(*Cell).Car()
}

// String ...
func (cell *Cell) String() string {
	switch {
	case cell.IsNull():
		return "()"
	case cell.IsList():
		return fmt.Sprintf("(%v %v)", cell.car, cell.cdr)
	default:
		return fmt.Sprintf("(%v . %v)", cell.car, cell.cdr)
	}
}

// IsNull ...
func (cell *Cell) IsNull() bool {
	return cell.car == nil && cell.cdr == nil
}

// IsList ...
func (cell *Cell) IsList() bool {
	if cell.IsNull() {
		return true
	}
	switch cell.cdr.(type) {
	case *Cell:
		return cell.cdr.(*Cell).IsList()
	default:
		return false
	}
}

// IsPair ...
func (cell *Cell) IsPair() bool {
	return !cell.IsNull()
}

// Len ...
func (cell *Cell) Len() int {
	if cell.IsNull() {
		return 0
	}
	return 1 + cell.cdr.(*Cell).Len()
}

// PushBack ...
func (cell *Cell) PushBack(exp Expression) {
	if cell.IsNull() {
		cell.car = exp
		cell.cdr = &Cell{}
	} else {
		cell.cdr.(*Cell).PushBack(exp)
	}
}

// Copy ...
func Copy(cell *Cell) *Cell {
	res := &Cell{}
	if cell.IsNull() {
		return res
	}
	car, ok := cell.car.(*Cell)
	if ok {
		res.car = Copy(car)
	} else {
		res.car = cell.car
	}
	cdr, ok := cell.cdr.(*Cell)
	if ok {
		res.cdr = Copy(cdr)
	} else {
		res.cdr = cell.cdr
	}
	return res
}
