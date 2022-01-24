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

func (b *Book) print() {
	fmt.Printf("Title: %-20s Price: %.2f\n", b.Title, b.Price)
}
