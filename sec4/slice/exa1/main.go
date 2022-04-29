// Slice example
package main

import "fmt"

func main() {
	langs := []string{"JavaScript", "Python", "HTML", "CSS", "Java", "SQL", "NoSQL", "C#", "Go"}

	printSlice("Top Programing Language", langs)

	// Adding some more to the list
	langs = append(langs, "Rust", "Perl")

	printSlice("Top Programing Language List 2", langs)

	// show the first Language
	fmt.Println()
	fmt.Println(langs[0], "is one of the highest demand Languages")
	fmt.Println(langs[1], "is the second best in the group")

	langs = langs[2:]
	fmt.Println("Print the rest of the Languages on the list", langs)

}

func printSlice(head string, slice []string) {
	fmt.Println()
	fmt.Println("-----", head, "-----")
	for i := 0; i < len(slice); i++ {
		el := slice[i]
		fmt.Println(el)
	}
}
