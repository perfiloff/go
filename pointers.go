// Реализуйте функцию CopyParent(p *Parent) Parent, которая создает копию структуры Parent и возвращает ее:
package main

import "fmt"

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func main() {
	cp := CopyParent(nil) // Parent{}
	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
			// {
			// 	Name: "Vasya",
			// 	Age:  22,
			// },
		},
	}
	cp = CopyParent(p)

	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{}
	fmt.Println(p.Children) // [{Andy 18}]
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	CopyPar := Parent{}

	CopyPar.Name = p.Name
	for _, child := range p.Children {
		CopyPar.Children = append(CopyPar.Children, Child{Name: child.Name, Age: child.Age})
	}

	return CopyPar
}

// func CopyParent(p *Parent) Parent {
// 	if p == nil {
// 		return Parent{}
// 	}

// 	cp := *p

// 	cp.Children = make([]Child, len(p.Children))
// 	copy(cp.Children, p.Children)

// 	return cp
// }
