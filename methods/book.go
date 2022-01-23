package main

import (
	"fmt"
	"time"
)

type Book struct {
	ISBN      string
	Title     string
	Price     float32
	Author    string
	PubDate   time.Time
	Publisher string
}

type ListBooks []Book

// PrintItem method print a single book with any idividual field
func (b *Book) PrintItem() {
	fmt.Printf("Book Title: %-20s Price %.2f\n", b.Title, b.Price)
}

// PrintList method print a slice of books
func (l *ListBooks) PrintList() {
	if len(*l) > 0 {
		for i, b := range *l {
			fmt.Printf("#%d Title: %-20s Price: %.2f\n", i, b.Title, b.Price)
		}
	} else {
		fmt.Println("The is no book in the list!")
	}
}
