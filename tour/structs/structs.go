package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structOne() {
	fmt.Println("WELCOME TO STRUCTS")
	v := Vertex{1, 2}
	v.X = 3
	fmt.Println(v)
	fmt.Println(v.X)
}

func main() {
	structOne()
}
