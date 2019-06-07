package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign k-v
	// emails["Bob"] = "bob@gmail.com"
	// emails["Gike"] = "gike@gmail.com"
	// emails["Jon"] = "jon@gmail.com"

	// Declare map and add k-v

	emails := map[string]string{"Bob": "bob@gmail.com", "Mark": "mark@gmail.com"}

	fmt.Println(emails)
	delete(emails, "Gike")
	fmt.Println(emails)
}
