// Pointer Example
package main

import "fmt"

type Counter struct {
	hits int
}

func increment(count *Counter) {
	count.hits += 1
	fmt.Println("Counter", count)
}

func replaceString(old *string, new string, count *Counter) {
	*old = new
	increment(count)
}

func main() {

	// make a copy of Counter struct
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replaceString(&hello, "Hi", &counter)

	fmt.Println(hello, world)

	str := []string{"Welcome", "Gophers", "Go", "Amazing", "Programming", "Language"}

	for _, s := range str {
		replaceString(&hello, s, &counter)
		fmt.Println(hello, world)
	}
}
