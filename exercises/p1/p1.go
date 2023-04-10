package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const appTitle = "Book Listing"
const separator = "---------------------------------"
const bookTitle = "Please enter book title: "
const bookDesc = "Please enter book description: "
const bookID = "Please enter book ID: "
const bookPrice = "Please enter book price: "

// Declare Book Struct data type
type Book struct {
	ID          string
	Title       string
	Description string
	Price       float64
}

// NewBook create a new book object
func NewBook(id, title, desc string, price float64) *Book {
	newBook := Book{
		ID:          id,
		Title:       title,
		Description: desc,
		Price:       price,
	}
	return &newBook
}

// PrintBook print out the book field's value
func (b *Book) PrintBook() {
	fmt.Printf("Book Title: %s\n", b.Title)
	fmt.Printf("Description: %s\n", b.Description)
	fmt.Printf("ID: %s\n", b.ID)
	fmt.Printf("Price: $%.2f\n", b.Price)
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

// Print App Info
func getBook() *Book {
	// Read user input
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println(appTitle)
	fmt.Println(separator)

	title := readUserInput(reader, bookTitle)
	desc := readUserInput(reader, bookDesc)
	id := readUserInput(reader, bookID)
	priceInput := readUserInput(reader, bookPrice)

	price, _ := strconv.ParseFloat(priceInput, 64)

	book := NewBook(id, title, desc, price)

	return book
}

func (b *Book) writeToFile() {
	file, _ := os.Create(b.ID + ".txt")

	content := fmt.Sprintf("Book Title: %s\nDescription: %s\nID: %s\nPrice: %.2f\n",
		b.Title, b.Description, b.ID, b.Price)

	file.WriteString(content)
	file.Close()
}

func main() {

	// Create an intance of Book data type
	harryPotter := Book{
		ID:          "0001",
		Title:       "Harry Potter Box Set",
		Description: "Harry Potter is a series of seven fantasy novels written by British author J. K.",
		Price:       59.99,
	}

	// Print the book data
	harryPotter.PrintBook()

	// Declare new bootk
	var theGivingTree *Book

	// create a new book using newBook function
	theGivingTree = NewBook("0002", "The Giving Tree", "The Giving Tree by Shele Silvertein", 10.75)

	// PrintBook
	theGivingTree.PrintBook()

	// getBook get book's info from user
	userInputBook := getBook()

	// Print Book Info
	userInputBook.PrintBook()

	// Prite book info to file
	userInputBook.writeToFile()
}

// # Your Tasks

// ## 1) Create a new type / data structure for storing product data (e.g. about a book)
// The data structure should contain these fields:
// 	- ID
// 	- Title / Name
// 	- Short description
// 	- price (number without currency, we'll just assume USD)
// ## 2) Create concrete instances of this data type in two different ways:
// 	- Directly inside of the main function
// 	- Inside of main, by using a "creation helper function"
// 	(use any values for title etc. of your choice)
// 	Output (print) the created data structures in the command line (in the main function)
// ## 3) Add a "connected function" that outputs the data + call that function from inside "main"
// ## 4) Change the program to fetch user input values for the different data fields
// 	and create only one concrete instance with that entered data.
// 	Output that instance data (via the connected function) then.
// ## 5) Bonus: Add a connected "store" function that writes that data into a file
// 	The filename should be the unique id, the function should be called at the
// 	end of main.
