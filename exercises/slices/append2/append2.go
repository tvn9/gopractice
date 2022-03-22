package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		pizza      []string
		depart     []time.Time
		graduation []int
		lights     []bool
	)

	now := time.Now()
	pizza = append(pizza, "Pepperoni", "Onions", "Extra cheese")
	depart = append(depart, now, now, now)
	graduation = append(graduation, 1970, 1999, 2022)
	lights = append(lights, true, false, true)

	fmt.Printf("Pizza topping: %s\n", pizza)
	fmt.Printf("Departure time: %s\n", depart)
	fmt.Printf("graduation year: %d\n", graduation)
	fmt.Printf("lights on/off: %t\n", lights)

}
