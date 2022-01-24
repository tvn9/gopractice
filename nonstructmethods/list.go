package main

type list []*Book

func (l list) print() {
	for _, it := range l {
		it.print()
	}
}
