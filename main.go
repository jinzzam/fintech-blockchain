package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	a := make([]int, 5)
	a = []int{1, 2, 3, 4, 5}
	printSlice("a", a)
	b := a[:2]
	printSlice("b", b)
	c := a[2:5]
	printSlice("c", c)
	d := a[:1]
	d = append(d)
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s %v\n", s, x)
}
