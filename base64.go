package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Ali Fahrial Anwar"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))

	fmt.Printf(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
