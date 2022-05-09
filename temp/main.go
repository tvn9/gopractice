// Interface
package main

import "fmt"

type MyInterface interface {
	Function1()
	Function2(i int) int
}

type MyType int

func execute(i MyInterface) {
	i.Function1()
}

func (m *MyType) Function1() {
	*m = *m * 8
	fmt.Println(*m)
}

func (m *MyType) Function2(i int) int {
	return i + i
}

func main() {

	m := MyType(10)

	execute(&m)

	temp := m.Function2(5)
	fmt.Println(temp)

}
