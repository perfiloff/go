// Реализуйте функцию MergeNumberLists(numberLists ...[]int) []int, которая принимает вариативный список слайсов чисел и объединяет их в 1, сохраняя последовательность:
// MergeNumberLists([]int{1, 2}, []int{3}, []int{4}) // [1, 2, 3, 4]

package main

import "fmt"

func main() {
	fmt.Println(MergeNumberLists([]int{1, 2}, []int{3}, []int{4}))
}

func MergeNumberLists(numberLists ...[]int) []int {
	res := make([]int, 0)
	for _, i := range numberLists {
		res = append(res, i...)
	}
	return res
}

// func MergeNumberLists(numberLists ...[]int) []int {
// 	mergedCap := 0
// 	for i := 0; i < len(numberLists); i++ {
// 		mergedCap += len(numberLists[i])
// 	}

// 	merged := make([]int, 0, mergedCap)
// 	for _, nl := range numberLists {
// 		merged = append(merged, nl...)
// 	}

// 	return merged
// }
