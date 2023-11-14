// Creation and Implemetation of Interface
package main

type MyInterface interface {
	Function1()
	Function2(n int) int
}

type MyType int

func (m MyType) Function1() {}

func (m MyType) Function2(n int) int {
	return n * n
}

func execute(i MyInterface) {
	i.Function1()
}
