// Реализуйте функции nextASCII(b byte) byte и prevASCII(b byte) byte, которые возвращают следующий или предыдущий символ ASCII таблицы соответственно. Например:
// nextASCII(byte('a')) // 'b'
// prevASCII(byte('b')) // 'a'
// Допускаем, что в функцию prevASCII не может прийти нулевой символ, а в функцию nextASCII — последний символ ASCII таблицы.

package main

import "fmt"

func main() {
	fmt.Println(string(nextASCII(byte('a'))))
	fmt.Println(prevASCII(byte('b')))
}

func nextASCII(b byte) byte {
	b++
	return b

}

func prevASCII(b byte) byte {
	b--
	return b
}
