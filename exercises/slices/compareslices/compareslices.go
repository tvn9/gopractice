package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesSlice := strings.Split(namesA, ", ")
	sort.Strings(namesSlice)

	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sort.Strings(namesB)

	fmt.Println("NamesSlice compare to NamesB Slice")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("namesSlice len = %d; namesB len = %d\n", len(namesSlice), len(namesB))
	if len(namesSlice) == len(namesB) {
		for i, a := range namesSlice {
			if a == namesB[i] {
				fmt.Printf("#%d %-10s = %s\n", i, a, namesB[i])
			}
		}
		fmt.Println("They are equal.")
	}
}
