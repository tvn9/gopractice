//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func lineReader(lines []string, lineCallBack func(line string)) {
	for _, l := range lines {
		lineCallBack(l)
	}
}

func main() {
	// lines is a slice of strings
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	// Define count holder variables
	var digitCnt, letterCnt, spaceCnt, puncCnt int

	// Define closure function
	lineFn := func(line string) {
		for _, l := range line {
			if unicode.IsDigit(l) {
				digitCnt += 1
			} else if unicode.IsSpace(l) {
				spaceCnt += 1
			} else if unicode.IsLetter(l) {
				letterCnt += 1
			} else if unicode.IsPunct(l) {
				puncCnt += 1
			} else {
				fmt.Println("now sure what it is", l)
			}
		}
	}

	// lineReader calls and pass lines, and lineFn to lineReader function
	lineReader(lines, lineFn)

	fmt.Printf("%d Digits\n", digitCnt)
	fmt.Printf("%d Letters\n", letterCnt)
	fmt.Printf("%d Spaces\n", spaceCnt)
	fmt.Printf("%d Punctuations\n", puncCnt)
}
