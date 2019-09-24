package main

import "fmt"

func main() {
	// Define Map
	emails := make(map[string]string)

	// Assign key values

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails["Bob"])

	// Delete
	delete(emails, "Bob")

	//declare
	// emails := map[string]string{"Bob": "Bob@gmail.com", "Mike": "Mike@gmail.com", "Sharon": "sharong@gmail.com"}
	fmt.Println(emails)
}
