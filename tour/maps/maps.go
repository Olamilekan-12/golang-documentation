package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m map[string]Vertex

func mapOne() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)

}

func main() {
	mapOne()
}
