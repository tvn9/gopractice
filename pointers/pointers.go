package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)
	fmt.Println("Address of age is", &age)

	myAge := &age

	fmt.Println("&age is assigned to myAge, therefore myAge is now holding the address of age", myAge)

	fmt.Println("To get the value of from age we can use dereferencing *myAge", *myAge)

	// Let try to see if we can change the value from age
	*myAge = age + 3

	fmt.Println("age after *myAge = 35, age is now", age)
	fmt.Println("Address of age", myAge)
	fmt.Println("Address of myAge is", &myAge)

	// Let try to change the value from myAge
	// myAge = 88 // Error because 88 is type int while myAge is type pointer *myAge

	// Now let try to pass age to function
	doubleNum := double(age)
	fmt.Println("doubleNum is = ", doubleNum) // => 70
	fmt.Println("age is now = ", age)         // => 35 value of age is still 35

	// Now let try to change the value of age from a function
	twoTimeAge := doubleAge(myAge)
	fmt.Println("twoTimeAge is =", twoTimeAge)
	fmt.Println("age is now =", age)
}

func double(num int) int {
	return num * 2
}

func doubleAge(ageAddr *int) int {
	*ageAddr = *ageAddr * 2
	return *ageAddr
}
