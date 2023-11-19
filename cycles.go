// Реализуйте функцию func Map(strs []string, mapFunc func(s string) string) []string, которая преобразует каждый элемент слайса strs с помощью функции mapFunc и возвращает новый слайс. Учтите, что исходный слайс, который передается как strs, не должен измениться в процессе выполнения.
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"Alex", "John", "Jeff"}
	fmt.Println(s)
	fmt.Println(Map(s, mapFunc))
}

func Map(strs []string, mapFunc func(s string) string) []string {
	newStrs := make([]string, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		newStrs = append(newStrs, mapFunc(strs[i]))
	}
	return newStrs
}

func mapFunc(s string) string {
	return strings.ToLower(s)
}

// func Map(strs []string, mapFunc func(s string) string) []string {
// 	mapped := make([]string, len(strs))
// 	for i, s := range strs {
// 		mapped[i] = mapFunc(s)
// 	}

// 	return mapped
// }
