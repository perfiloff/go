// Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64, которая возвращает отсортированный слайс, состоящий из уникальных идентификаторов userIDs. Обработка должна происходить in-place, то есть без выделения доп. памяти.
package main

import (
	"fmt"
	"sort"
)

func main() {
	userIDs := []int64{1, 5, 7, 4, 1, 6, 6, 6, 7, 8, 9, 12, 12}

	fmt.Println(UniqueSortedUserIDs(userIDs))
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) <= 1 {
		return userIDs
	}
	sort.Slice(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	// fmt.Println(userIDs)
	pointer := 0
	for i := 1; i < len(userIDs); i++ {
		// fmt.Printf("Pointer: (%d, %d), i: (%d, %d),\n", pointer, userIDs[pointer], i, userIDs[i])
		if userIDs[pointer] != userIDs[i] {

			pointer++
			userIDs[pointer] = userIDs[i]
		}
	}
	return userIDs[:pointer+1]
}
