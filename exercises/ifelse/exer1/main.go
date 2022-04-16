// Validate User/Password Challenge
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = `
Please enter username and password
Example: command [username] [password]`
	userID, passID = "chris", "8888"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		fmt.Println()
		return
	}

	user, pass := args[0], args[1]

	user = strings.ToLower(user)
	pass = strings.ToLower(pass)

	if user != userID || pass != passID {
		fmt.Println("username of password does not match!")
		return
	} else if user == userID && pass == passID {
		fmt.Println("Access granted!")
	}
}
