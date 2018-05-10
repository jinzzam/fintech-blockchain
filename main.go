package main

import (
	"fmt"
	"math"
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

	fmt.Println(pow(3, 3, 20), pow(3, 2, 10))

}

func printSlice(s string, x []int) {
	fmt.Printf("%s %v\n", s, x)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
