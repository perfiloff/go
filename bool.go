package main

import (
	"fmt"
)

func main() {
	a := IsValid(-1, "Hello, world")
	fmt.Println(a)
}

func IsValid(id int, text string) bool {
	if id <= 0 || text == "" {
		return false
	} else {
		return true
	}
}
