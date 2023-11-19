// Реализуйте интерфейс IVoiceable для структур Cat, Cow и Dog так, чтобы при вызове метода Voice экземпляр структуры Cat возвращал строку "Мяу", экземпляр Cow строку "Мууу", а экземпляр Dog сообщение Гав:
package main

import "fmt"

type IVoiceable interface {
	Voice() string
}

type Cat struct {
	// …
}

type Cow struct {
	// …
}

type Dog struct {
	// …
}

func (ct Cat) Voice() string {
	return "Мяу"
}
func (cw Cow) Voice() string {
	return "Мууу"
}
func (dg Dog) Voice() string {
	return "Гав"
}

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}
