// Реализуйте методы структуры Counter, представляющую собой счётчик, хранящий неотрицательное целочисленное значение и позволяющий это значение изменять:
//     метод Inc(delta int) должен увеличивать текущее значение на delta единиц (на 1 по умолчанию),
//     метод Dec(delta int) должен уменьшать текущее значение на delta единиц.
// c := Counter{}
// c.Inc(0)
// c.Inc(0)
// c.Inc(40)
// c.Value // 42

// c.Dec(0)
// c.Dec(30)
// c.Value // 11

// c.Dec(100)
// c.Value // 0

package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value += delta
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		delta = 1
	}
	if c.Value >= delta {
		c.Value -= delta
	} else if c.Value < delta {
		c.Value = 0
	}
}

func main() {
	c := Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	fmt.Println(c.Value) // 42

	c.Dec(0)
	c.Dec(30)
	fmt.Println(c.Value) // 11

	c.Dec(100)
	fmt.Println(c.Value) // 0
}

// func Max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// // BEGIN
// func (c *Counter) Inc(delta int) {
// 	if delta == 0 {
// 		delta = 1
// 	}
// 	c.Value = Max(c.Value+delta, 0)
// }

// func (c *Counter) Dec(delta int) {
// 	if delta == 0 {
// 		delta = 1
// 	}
// 	c.Inc(-delta)
// }
