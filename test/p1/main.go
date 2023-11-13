package main

import (
	"regexp"
)

func IsValidEmail(addr string) bool {
	re, err := regexp.Compile(`.+@.+\..+`)
	if err != nil {
		panic("failed to complile regex")
	} else {
		return re.Match([]byte(addr))
	}
}

func main() {

}
