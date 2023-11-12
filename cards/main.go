package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, c := range cards {
		fmt.Println(i, c)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
