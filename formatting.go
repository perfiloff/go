// Реализуйте функцию generateSelfStory(name string, age int, money float64) string, которая генерирует предложение, подставляя переданные данные в возвращаемую строку. Например:
// generateSelfStory("Vlad", 25, 10.00000025) // "Hello! My name is Vlad. I'm 25 y.o. And I also have $10.00 in my wallet right now."
// Шаблон возвращаемой строки: Hello! My name is [name]. I'm [age] y.o. And I also have $[money with precision 2] in my wallet right now.

package main

import "fmt"

func main() {
	fmt.Println(generateSelfStory("Vlad", 25, 10.00000025))
}

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}
