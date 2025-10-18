package main

import "fmt"

func slicesOne() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func slicesTwo() {
	var values []int

	for num := 0; num <= 10; num++ {
		values = append(values, num)
	}
	fmt.Println(values)
}

func slicesThree() {
	newValues := []int{
		1, 2, 3, 4,
	}
	fmt.Println(newValues)
	newValues[0] = 8
	fmt.Println(newValues)
}

func main() {
	slicesOne()
	slicesTwo()
	slicesThree()
}
