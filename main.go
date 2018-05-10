package main

import (
	"fmt"
	"math"
)

const ( //상수 : 문자, 문자열, boolean, 숫자 타입 중 하나
	//타입을 지정해주지 않은 상수는 문맥에 따라 타입을 자동으로 가짐
	v = 122
	w = v + 4
)

type Vertex struct { //구조체 (필드(데이터)들의 조합)
	//type선언으로 struct의 이름을 저장할 수 있음
	X int
	Y int
}

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

	var x, y, z int = 1, 2, 3             /*변수 선언*/
	C, python, java := true, false, "no!" //:=를 사용하면 var과 명시적입 타입 생략 가능
	fmt.Println(x, y, z, C, python, java)
	fmt.Println(v, w)

	fmt.Println(pow(3, 3, 20), pow(3, 2, 10))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("0~10 sum : ", sum)

	p := Vertex{1, 2}
	q := &p
	q.X = 19
	q.Y = 11
	fmt.Println("구조체를 써보자 : ", p)
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
