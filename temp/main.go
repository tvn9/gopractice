package main

import "fmt"

func main() {
	var books [5]string
	games := []string{"Little Plannet", "Call of duty", "minescrapt"}

	newGames := games

	// the only way to compare slices is to compare the slice's elements
	for i, g := range games {
		for j, ng := range newGames {
			if g == ng {
				fmt.Printf("%d %v is equal to %d %v\n", i, g, j, ng)
			}
		}
	}

	// This will not work for slice, Go does not allow compare slice by == sign
	/*
		if games == newGames {
			fmt.Println(games, "and", newGames, "are equal!")
		}
	*/

	// but we can compare the len of the slices as follow
	if len(games) == len(newGames) {
		fmt.Println("they are equal")
	}

	fmt.Printf("books %T\n", books)
	fmt.Printf("books %d\n", len(books))

	fmt.Printf("games %T\n", games)
	fmt.Printf("games %d\n", len(games))
	fmt.Printf("games %t\n", games == nil)

	books[0] = "The Go Programming Language"
	books[1] = "The Practice of Programming"
	books[2] = "Introducing Go"

	fmt.Printf("books: %q\n", books)
	fmt.Printf("games: %q\n", games)
}
