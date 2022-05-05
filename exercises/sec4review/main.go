//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

// Define some customer types
type Title string
type Name string

// LendAudit struct keeping variable to checkOut/checkIn time
type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

// Member struct holds member name and a map of book title lendAudit time
type Member struct {
	name  Name
	books map[Title]LendAudit
}

// BookEntry struct holds total number of books in store and total
// number of books being lended to members
type BookEntry struct {
	total  int // total books in Liberal
	lended int // total lended
}

// Library struct holds map of member's names and book's title and audit
// records
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

// printMemberAudit function prints the individual member and book audits
// information
func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

// printMemberAudits function prints a list of Library's members book audits info
func printMemberAudits(Library *Library) {
	for _, member := range Library.members {
		printMemberAudit(&member)
	}
}

// printLibraryBooks function prints the Library's book title and lended status
func printLibraryBooks(Library *Library) {
	fmt.Println()
	for title, book := range Library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

// returnBook function process the book return from library's members
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out the book")
		return true
	}
	// Update library
	book.lended -= -1
	library.books[title] = book
	// Update member information
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

// checkoutBook process the book checkout and update total count of book lended.
func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("The book is out of stock")
		return false
	}

	// Update library
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

// main function is start program execution
func main() {
	// Define a library object that holds a map of book title, bookEntry record
	// and a map of member's name as key to Member object
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	// Creating a list of books, with total count, and lended count default to 0.
	library.books["Webapps in Go"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.books["Let's Learn Go"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.books["Go Bootcampt"] = BookEntry{
		total:  5,
		lended: 0,
	}

	// Insert members to library's map of Name and Book list
	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susanna"] = Member{"Susanna", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
