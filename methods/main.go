package main

func main() {

	// Defind some books
	book1 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Price: price{
			price:     10,
			priceUp:   0.0,
			priceDn:   .2,
			salePrice: 0.0,
		},
	}
	book1.Price.salePrice = book1.discount(&book1.Price.priceDn)

	book2 := Book{
		Title:  "1984",
		Author: "George Orwell",
		Price: price{
			price:   38.95,
			priceUp: 0.0,
			priceDn: .5,
		},
	}
	book2.Price.salePrice = book2.discount(&book2.Price.priceDn)

	book3 := Book{
		Title:  "War of Peace",
		Author: "Leo Tolstoy",
		Price: price{
			price:   28.95,
			priceUp: .2,
			priceDn: 0,
		},
	}
	book3.Price.salePrice = book3.increase(&book3.Price.priceUp)

	books := ListBooks{&book1, &book2, &book3}

	// Print one book at a time
	book1.PrintItem()
	book2.PrintItem()
	book3.PrintItem()

	// Print a list of books
	books.PrintList()
}
