package main

import (
	"fmt"
)

// TODO: Change rescursion to for loop

// List ...
type List struct {
	car Expression
	cdr Expression
}

// Car ...
func (list *List) Car() (Expression, error) {
	if list.IsNull() {
		return nil, fmt.Errorf("attempt to apply car on nil %v", list)
	}
	return list.car, nil
}

// Cdr ...
func (list *List) Cdr() (Expression, error) {
	if list.IsNull() {
		return nil, fmt.Errorf("attempt to apply cdr on nil %v", list)
	}
	return list.cdr, nil
}

// Cadr second
func (list *List) Cadr() (Expression, error) {
	cdr, err := list.Cdr()
	if err != nil {
		return nil, err
	}
	return cdr.(*List).Car()
}

// Cddr ...
func (list *List) Cddr() (Expression, error) {
	cdr, err := list.Cdr()
	if err != nil {
		return nil, err
	}
	return cdr.(*List).Cdr()
}

// Cdddr ...
func (list *List) Cdddr() (Expression, error) {
	cdr, err := list.Cddr()
	if err != nil {
		return nil, err
	}
	return cdr.(*List).Cdr()
}

// Caddr third
func (list *List) Caddr() (Expression, error) {
	cddr, err := list.Cddr()
	if err != nil {
		return nil, err
	}
	return cddr.(*List).Car()
}

// IsNull ...
func (list *List) IsNull() bool {
	return list.car == nil && list.cdr == nil
}

// IsList ...
func (list *List) IsList() bool {
	if list.IsNull() {
		return true
	}
	switch list.cdr.(type) {
	case *List:
		return list.cdr.(*List).IsList()
	default:
		return false
	}
}

// IsPair ...
func (list *List) IsPair() bool {
	return !list.IsNull()
}

// Len ...
func (list *List) Len() int {
	if list.IsNull() {
		return 0
	}
	return 1 + list.cdr.(*List).Len()
}

// PushBack ...
func (list *List) PushBack(exp Expression) {
	if list.IsNull() {
		list.car = exp
		list.cdr = &List{}
	} else {
		list.cdr.(*List).PushBack(exp)
	}
}

// MakeList ...
func MakeList(args ...Expression) *List {
	list := &List{}
	for i := len(args) - 1; i >= 0; i-- {
		list = &List{args[i], list}
	}
	return list
}

// Copy ...
func Copy(list *List) *List {
	res := &List{}
	if list.IsNull() {
		return res
	}
	car, ok := list.car.(*List)
	if ok {
		res.car = Copy(car)
	} else {
		res.car = list.car
	}
	cdr, ok := list.cdr.(*List)
	if ok {
		res.cdr = Copy(cdr)
	} else {
		res.cdr = list.cdr
	}
	return res
}
