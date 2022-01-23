package main

func main() {
	// Defind some books
	book1 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Price:  35.95,
	}
	book2 := Book{
		Title:  "1984",
		Author: "George Orwell",
		Price:  38.95,
	}
	book3 := Book{
		Title:  "War of Peace",
		Author: "Leo Tolstoy",
		Price:  28.95,
	}

	books := ListBooks{book1, book2, book3}

	// Print one book at a time
	book1.PrintItem()
	book2.PrintItem()
	book3.PrintItem()

	// Print a list of books
	books.PrintList()
}
