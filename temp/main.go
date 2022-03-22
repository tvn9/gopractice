package main

import "fmt"

func main() {
	games := [5]string{"Little Plannet", "Call of duty", "minescrapt"}

	games[3] = "candy crush"
	games[4] = "Roblox"

	fmt.Println(games, len(games))

	first3 := games[:3]
	fmt.Printf("First 3 [:3] = %q\n", first3)

	l := len(games)
	last2 := games[l-2:]
	fmt.Printf("Last 2 [l-2:] - %q\n", last2)

	games[0] = "Guita herro"

	fmt.Println(games)

	fmt.Printf("First 3 [:3] = %q\n", first3)

}
