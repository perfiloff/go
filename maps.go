// Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс, состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.
package main

import "fmt"

func main() {
	userIDs := []int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}

	fmt.Println(UniqueUserIDs(userIDs))
}

func UniqueUserIDs(userIDs []int64) []int64 {
	cache := make(map[int64]struct{})
	newSlice := make([]int64, 0)
	var ok bool

	for i := 0; i < len(userIDs); i++ {
		_, ok = cache[userIDs[i]]
		cache[userIDs[i]] = struct{}{}
		if ok != true {
			newSlice = append(newSlice, userIDs[i])
		}
	}
	return newSlice

}

// func UniqueUserIDs(userIDs []int64) []int64 {
// 	// пустая структура struct{} — это тип данных, который занимает 0 байт
// 	// используется, когда нужно проверять в мапе только наличие ключа
// 	processed := make(map[int64]struct{})

// 	uniqUserIDs := make([]int64, 0)
// 	for _, uid := range userIDs {
// 		_, is_exists := processed[uid]
// 		if is_exists {
// 			continue
// 		}

// 		uniqUserIDs = append(uniqUserIDs, uid)
// 		processed[uid] = struct{}{}
// 	}

// 	return uniqUserIDs
// }
