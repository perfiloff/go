// В Go нет встроенной функции удаления элемента из слайса. Реализуйте функцию Remove(nums []int, i int) []int, которая удаляет элемент по индексу i из слайса nums. Если приходит несуществующий индекс, то из функции возвращается исходный слайс. Порядок элементов может быть нарушен после удаления элемента.

package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	n := make([]int, 0, 4)
	n = Remove(s, 2)
	fmt.Println(n)

}

func Remove(nums []int, i int) []int {
	var new int
	if len(nums) < i+1 || i < 0 {
		return nums
	}
	new = len(nums) - 1
	nums[i] = nums[new]
	return nums[:new]
}
