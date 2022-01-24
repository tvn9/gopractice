package main

import (
	"fmt"
	"time"
)

type price struct {
	price     float32
	priceDn   float32
	priceUp   float32
	salePrice float32
}

type Book struct {
	ISBN      string
	Title     string
	Price     price
	Author    string
	PubDate   time.Time
	Publisher string
}

type ListBooks []*Book

// PrintItem method print a single book with any idividual field
func (b *Book) PrintItem() {
	fmt.Printf("Book Title: %-20s Price %.2f\n", b.Title, b.Price.price)
}

// PrintList method print a slice of books
func (l *ListBooks) PrintList() {
	if len(*l) > 0 {
		for i, b := range *l {
			fmt.Printf("#%d Title: %-20s Price: %.2f, Discount: %.2f, Inscrease: %.2f, Sale Price: %.2f\n", i, b.Title, b.Price.price, b.Price.priceDn, b.Price.priceUp, b.Price.salePrice)
		}
	} else {
		fmt.Println("The is no book in the list!")
	}
}

// discount
func (p *Book) discount(d float32) float32 {
	p.Price.salePrice = p.Price.price * (1 - d)
	return p.Price.salePrice
}

// increase
func (p *Book) increase(d float32) float32 {
	p.Price.salePrice = p.Price.price * (1 + d)
	return p.Price.salePrice
}
