package main

import (
	"fmt"
	"strings"
)

// // когда передается mode "dash", нужно заменить все пробелы на дефисы "-"
// ModifySpaces("hello world", "dash") // hello-world

// // когда передается mode "underscore", нужно заменить все пробелы на нижние подчеркивания "_"
// ModifySpaces("hello world", "underscore") // hello_world

// // когда передается неизвестный или пустой mode, заменяем все пробелы на звездочки "*"
// ModifySpaces("hello world", "unknown") // hello*world
// ModifySpaces("hello world", "") // hello*world

func main() {
	fmt.Println(ModifySpaces(" great ", ""))
}

func ModifySpaces(s, mode string) string {
	var new_char string

	switch mode {
	default:
		new_char = "*"
	case "dash":
		new_char = "-"
	case "underscore":
		new_char = "_"
	}

	return strings.ReplaceAll(s, " ", new_char)
}
