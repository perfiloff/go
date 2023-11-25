// Реализуйте функцию isASCII(s string) bool, которая возвращает true, если строка s состоит только из ASCII символов.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isASCII("hello \U0001F970"))
}

func isASCII(s string) bool {
	for _, ch := range s {
		fmt.Println(ch)
		if int(ch) > 127 {
			return false
		}
	}
	return true
}

// func isASCII(s string) bool {
// 	for _, r := range s {
// 		if r > unicode.MaxASCII {
// 			return false
// 		}
// 	}

// 	return true
// }
