package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Greetings("Man"))
}

func Greetings(name string) string {

	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)

	return fmt.Sprintf("Hello, %s", name)
}
