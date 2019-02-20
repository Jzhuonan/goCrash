package main

import "fmt"

func main() {

	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["1"] = "thedog@gmail.com"
	// emails["2"] = "jack@gmail.com"
	// emails["3"] = "tom@gmail.com"

	// Define map and assign kv
	emails := map[string]string{"1": "thedog@gmail.com", "2": "jack@gmail.com", "3": "tom@gmail.com"}

	// delete(emails, "2")

	fmt.Println(emails["3"])
	fmt.Println(emails)
}