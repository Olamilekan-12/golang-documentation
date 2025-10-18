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

func slicesFour() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		k bool
	}{
		{1, true},
		{2, false},
	}

	fmt.Println(s)
}

func slicesFive() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func main() {
	slicesOne()
	slicesTwo()
	slicesThree()
	slicesFour()
	slicesFive()
}
