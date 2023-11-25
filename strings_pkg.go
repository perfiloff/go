// В пакете unicode есть функция unicode.Is(unicode.Latin, rune), которая проверяет, что руна является латинским символом или нет.
// Реализуйте функцию latinLetters(s string) string, которая возвращает только латинские символы из строки s. Например:latinLetters("привет world!") // "world"

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(latinLetters("привет world!"))
}

func latinLetters(s string) string {
	sb := &strings.Builder{}
	for _, ch := range s {
		if unicode.Is(unicode.Latin, ch) {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}
