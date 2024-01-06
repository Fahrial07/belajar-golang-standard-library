package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[i].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	users := []User{
		{"Budi", 15},
		{"Ubed", 28},
		{"Ali", 25},
		{"Siti", 60},
		{"Andi", 28},
		{"Ari", 29},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
