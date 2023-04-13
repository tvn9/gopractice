package main

import "fmt"

func main() {

	hobbies := make([]string, 2, 10)
	fmt.Println(hobbies, len(hobbies), cap(hobbies))

	moreHobbies := new([]string)
	fmt.Println(moreHobbies, len(*moreHobbies), cap(*moreHobbies))

	hobbies[0] = "Tennis"
	hobbies[1] = "Socker"

	hobbies = append(hobbies, "Cooking", "Travelling")

	fmt.Println(hobbies, len(hobbies), cap(hobbies))

}
