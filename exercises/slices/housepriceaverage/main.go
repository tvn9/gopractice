package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)

	var (
		loc                        []string
		sizes, beds, baths, prices []int
	)

	hd := strings.Split(header, separator)
	// fmt.Printf("header: %q\n", hd)

	dataSlice := strings.Split(data, "\n")
	//	fmt.Printf("Data: %v\n", dataSlice)

	for _, h := range dataSlice {
		row := strings.Split(h, separator)
		loc = append(loc, row[0])

		size, _ := strconv.Atoi(row[1])
		sizes = append(sizes, size)

		bed, _ := strconv.Atoi(row[2])
		beds = append(beds, bed)

		bath, _ := strconv.Atoi(row[3])
		baths = append(baths, bath)

		price, _ := strconv.Atoi(row[4])
		prices = append(prices, price)
	}

	// Housing data table
	fmt.Println()
	for _, h := range hd {
		fmt.Printf("%-15s", h)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 70))

	// Print data
	for i := range loc {
		fmt.Printf("%-14s %-14d %-14d %-14d %d\n", loc[i], sizes[i], beds[i], baths[i], prices[i])
	}
	fmt.Println()

	// Housing price average
	fmt.Println(strings.Repeat("=", 70))

	var asize, abed, abath, aprice float64
	for i := range loc {
		asize += float64(sizes[i])
		abed += float64(beds[i])
		abath += float64(baths[i])
		aprice += float64(prices[i])
	}
	fmt.Printf("%-14s %-14.2f %-14.2f %-14.2f %.2f\n",
		"",
		asize/float64(len(sizes)),
		abed/float64(len(beds)),
		abath/float64(len(baths)),
		aprice/float64(len(prices)))

}
