// Реализуйте функцию SafeWrite(nums [5]int, i, val int) [5]int, которая записывает значение val в массив nums по индексу i, если индекс находится в рамках массива. В противном случае массив возвращается без изменения.
package main

import "fmt"

func main() {
	var my_num [5]int
	fmt.Println(SafeWrite(my_num, 0, 10))

}

func SafeWrite(nums [5]int, i, val int) [5]int {
	var arr_len int
	arr_len = len(nums)
	if i < arr_len && i >= 0 {
		nums[i] = val
	}
	return nums
}
