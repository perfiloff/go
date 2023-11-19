// Реализуйте функцию IntsCopy(src []int, maxLen int) []int, которая создает копию слайса src с длиной maxLen. Если maxLen равен нулю или отрицательный, то функция возвращает пустой слайс []int{}. Если maxLen больше длины src, то возвращается полная копия src.

package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	fmt.Println(IntsCopy(src, 2))
}

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	} else if maxLen > len(src) {
		return src
	}

	newSlice := make([]int, maxLen)
	copy(newSlice, src)

	return newSlice
}
