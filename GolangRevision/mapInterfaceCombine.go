package main

import "fmt"

func main() {
	// Define a map with string keys and interface{} values
	data := map[string]interface{}{
		"name":    "Santosh",
		"age":     42,
		"isAdmin": true,
		"tags":    []string{"go", "developer"},
	}

	// Accessing and printing the values
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Is Admin:", data["isAdmin"])
	fmt.Println("Tags:", data["tags"])
}
