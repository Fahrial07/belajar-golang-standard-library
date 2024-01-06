package main

import "fmt"

func main() {
	firstName := "Ali"
	lastName := "Fahrial"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
