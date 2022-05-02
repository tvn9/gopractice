// Map example
package main

import "fmt"

func main() {

	// Create a map variable
	cityPopulation := make(map[string]int)

	// insert map elements
	cityPopulation["New York"] = 8622257
	cityPopulation["Los Angeles"] = 4085014
	cityPopulation["Chicago"] = 2670406
	cityPopulation["Houston"] = 2378146
	cityPopulation["Phoenix"] = 1743469

	// Print out map
	fmt.Println(cityPopulation)
	fmt.Println(cityPopulation["Chicago"])

	// Different way to create map
	worlTopIncome := map[string]int{
		"Luxembourg":           118001,
		"Singapore":            97057,
		"Ireland":              94392,
		"Qata":                 93508,
		"Switzerland":          72874,
		"Norway":               65800,
		"United States":        63416,
		"Brunei":               62371,
		"Hong Kong":            59520,
		"Denmark":              58932,
		"United Arab Emirates": 58753,
	}

	// Print map
	fmt.Println(worlTopIncome)

	// Delete map elements
	delete(cityPopulation, "Phoenix")
	delete(worlTopIncome, "United Arab Emirates")

	fmt.Println(cityPopulation)
	fmt.Println(worlTopIncome)

	// Reading map data
	fmt.Println(cityPopulation["New York"])
	fmt.Println(worlTopIncome["New York"])

	// Check if element exist in map
	item, ok := worlTopIncome["Singapore"]
	if !ok {
		fmt.Println(item, "not exist")
	}
	fmt.Println("Singapore", item)

	// Map iteration
	fmt.Println("--------- Top City Population ---------")
	numOfPeople := 0
	for k, v := range cityPopulation {
		numOfPeople += v
		fmt.Printf("City %q - Population: %d\n", k, v)
	}
	fmt.Printf("Total combined most populated cities %d\n", numOfPeople)

	fmt.Println("--------- World Top GDP per Capita ---------")
	for k, v := range cityPopulation {
		fmt.Printf("City %q - income per capita: %d\n", k, v)
	}

}
