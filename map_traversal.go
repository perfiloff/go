// Реализуйте функцию MostPopularWord(words []string) string, которая возвращает самое часто встречаемое слово в слайсе. Если таких слов несколько, то возвращается первое из них.

package main

import "fmt"

func main() {
	test := []string{"Willy", "John", "Willy", "Anna", "John", "Peter", "John", "Willy"}
	fmt.Println(MostPopularWord(test))
}

// func MostPopularWord(words []string) string {
// 	wordsMap := make(map[string]int, 0)
// 	max := 0
// 	resWord := ""

// 	for _, word := range words {
// 		wordsMap[word]++
// 		if wordsMap[word] > max {
// 			max = wordsMap[word]
// 			resWord = word
// 		}
// 	}
// 	return resWord
// }

func MostPopularWord(words []string) string {
	wordsMap := make(map[string]int)
	// var qty int
	// var word string

	for _, word := range words {
		qty, is_exists := wordsMap[word]
		if is_exists {
			qty++
			wordsMap[word] = qty
		} else {
			wordsMap[word] = 1
		}
	}

	qty := 0
	var res string
	for _, word := range words {
		if wordsMap[word] > qty {
			res = word
			qty = wordsMap[word]
		}
	}
	return res
}
