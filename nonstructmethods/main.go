package main

func main() {

	book1 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Price:  35.95,
	}

	book2 := Book{
		Title:  "1984",
		Author: "George Orwell",
		Price:  20.99,
	}

	book3 := Book{
		Title:  "War of Peace",
		Author: "Leo Tolstoy",
		Price:  33.95,
	}

	var items []*Book

	items = append(items, &book1, &book2, &book3)

	// convert untype items to type *list => []*Book
	my := list(items)

	// after conversion variable my now has access to method print()
	my.print()
}
