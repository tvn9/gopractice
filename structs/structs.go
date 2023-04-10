package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName    string
	lastName     string
	birthdate    string
	job          string
	education    string
	sex          string
	merriage     string
	employment   string
	annualIncome int
	createdDate  time.Time
}

// NewUser function create new user
func NewUser(fName, lName, bDate, job, edu, sex, merr, emp string, annIn int) *User {
	created := time.Now()
	user := User{
		firstName:    fName,
		lastName:     lName,
		birthdate:    bDate,
		job:          job,
		education:    edu,
		sex:          sex,
		merriage:     merr,
		employment:   emp,
		annualIncome: annIn,
		createdDate:  created,
	}

	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	var newUser *User

	firstName := getUserInput("Please enter you first name: ")
	lastName := getUserInput("Please enter you last name: ")
	birthdate := getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	// Create new user by calling NewUser function
	newUser = NewUser(firstName, lastName, birthdate, "job", "edu", "sex", "merriage", "emp", 0)

	// Output user info
	newUser.outPutUserInfo()

}

func getUserInput(input string) string {

	fmt.Print(input)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput

}

func (u *User) outPutUserInfo() {
	fmt.Printf("First name: %s\n", u.firstName)
	fmt.Printf("Last name: %s\n", u.lastName)
	fmt.Printf("birth date: %s\n", u.birthdate)
}
