package main

import (
	"fmt"
)

// func slicesOne() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

// 	var s []int = primes[1:4]
// 	fmt.Println(s)
// }

// func slicesTwo() {
// 	var values []int

// 	for num := 0; num <= 10; num++ {
// 		values = append(values, num)
// 	}
// 	fmt.Println(values)
// }

// func slicesThree() {
// 	newValues := []int{
// 		1, 2, 3, 4,
// 	}
// 	fmt.Println(newValues)
// 	newValues[0] = 8
// 	fmt.Println(newValues)
// }

// func slicesFour() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q)

// 	r := []bool{true, false, true, true, false, true}
// 	fmt.Println(r)

// 	s := []struct {
// 		i int
// 		k bool
// 	}{
// 		{1, true},
// 		{2, false},
// 	}

// 	fmt.Println(s)
// }

// func slicesFive() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	s = s[1:4]
// 	fmt.Println(s)
// 	s = s[:2]
// 	fmt.Println(s)

// 	s = s[1:]
// 	fmt.Println(s)
// }

// func sliceSix() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)

// 	// Slice the slice to give it zero length.
// 	s = s[:0]
// 	printSlice(s)

// 	// Extend its length.
// 	s = s[:4]
// 	printSlice(s)

// 	// Drop its first two values.
// 	s = s[2:]
// 	printSlice(s)
// }

// func sliceSeven() {
// 	var s []int
// 	fmt.Println(s, len(s), cap(s))
// 	if s == nil {
// 		fmt.Println("nil!")
// 	}
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// func sliceEight() {
// 	a := make([]int, 5)
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c := b[:2]
// 	printSlice("c", c)

// 	d := c[2:5]
// 	printSlice("d", d)
// }

// func sliceNine() {
// 	board := [][]string{
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 	}

// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "0"
// 	board[0][2] = "X"

// 	for i := 0; i < len(board); i++ {
// 		fmt.Printf("%s\n", strings.Join(board[i], " "))
// 	}
// }

func sliceTen() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// slicesOne()
	// slicesTwo()
	// slicesThree()
	// slicesFour()
	// slicesFive()
	// sliceSix()
	// sliceSeven()
	// sliceEight()
	// sliceNine()
	sliceTen()
}
