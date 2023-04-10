package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://www.google.com",
		"amazon": "https://www.amazon.com",
	}

	// Print map
	fmt.Println(websites)
	// Look for specific key
	fmt.Println(websites["amazon"])

	// Adding new key value pair to map
	websites["LinkedIn"] = "https://www.linkedin.com"
	websites["facebook"] = "https://www.linkedin.com"
	fmt.Println(websites)

	// Overwrite a key/value
	websites["facebook"] = "https://www.facebook.com"
	fmt.Println(websites)

	// Delete a key/value pair
	delete(websites, "facebook")
	fmt.Println(websites)
}
