package main

import "fmt"

type Animal interface {
	GetInfo() string
}

type Cat struct {
	Name  string
	Color string
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("Cat: %s, Color: %s", c.Name, c.Color)
}

func printAnimal(animal Animal) {
	fmt.Println(animal.GetInfo())
}

func main() {
	cat := Cat{
		Name:  "test",
		Color: "Gray",
	}
	printAnimal(cat)
}
