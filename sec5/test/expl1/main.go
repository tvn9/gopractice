// Testing in Go example
package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

func init() {
	compliled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regular expression")
	}
	EmailExpr = compliled
	fmt.Println("regular expression compiled successfully")
}

func isValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@example.com"))
	fmt.Println(isValidEmail("invalid@example"))
}
