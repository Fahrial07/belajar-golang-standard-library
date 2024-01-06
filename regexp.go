package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])i`)

	fmt.Println(regex.MatchString("ali"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("eDi"))

	fmt.Println(regex.FindAllString("ali adi ada asi ani ami abi aci ari ati", 10))

}
