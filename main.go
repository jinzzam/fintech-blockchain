package main

import (
	"fmt"
)

const (
	v = 122
	w = v + 4
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

	var x, y, z int = 1, 2, 3
	C, python, java := true, false, "no!"
	fmt.Println(x, y, z, C, python, java)
	fmt.Println(v, w)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s %v\n", s, x)
}
