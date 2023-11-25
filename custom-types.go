// Представим, что есть структура Person, содержащая возраст человека:
// type Person struct {
//     Age uint8
// }
// Реализуйте тип PersonList (слайс структур Person), с методом (pl PersonList) GetAgePopularity() map[uint8]int, который возвращает мапу, где ключ — возраст, а значение — кол-во таких возрастов:
// pl := PersonList{
//   {Age: 18},
//   {Age: 44},
//   {Age: 18},
// }
// pl.GetAgePopularity() // map[18:2 44:1]

package main

import "fmt"

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	res := make(map[uint8]int, 0)
	for _, age := range pl {
		res[age.Age]++
	}

	return res
}

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}
	fmt.Println(pl.GetAgePopularity())
}
