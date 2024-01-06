package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Ali Fahrial Anwar", "Ali"))
	fmt.Println(strings.Split("Ali Fahrial Anwar", " "))
	fmt.Println(strings.ToLower("Ali Farial Anwar"))
	fmt.Println(strings.ToUpper("Ali Fahrial Anwar"))
	fmt.Println(strings.Trim("    Ali Fahrial Anwar    ", " "))
	fmt.Println(strings.ReplaceAll("Ali Fahrial Ali Anwar", "Ali", "Muhammad"))

}
