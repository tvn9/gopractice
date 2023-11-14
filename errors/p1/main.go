// Error Interface from Go standard library
/*
type error interface {
	Error() string
}
*/
package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no element at index %d", index)
	} else {
		return s.values[index], nil
	}

}

func divide(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("error: can not divide by zero")
	} else {
		return l / r, nil
	}
}

func main() {
	stuff := Stuff{}

	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// divide
	result, err := divide(9, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
