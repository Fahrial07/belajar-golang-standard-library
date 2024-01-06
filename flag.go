package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "Database username")
	var password *string = flag.String("password", "root", "Database password")
	var host *string = flag.String("host", "localhost", "Database host")
	var port *int = flag.Int("port", 0, "Database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)
}
