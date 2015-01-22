package main

import "fmt"

type point struct {
	x, y int
}

func NewPoint(x, y int) *point {
	return &point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	p := point{20, 40}
	fmt.Println("p:", p)

	p2 := point{x: 20, y: 40}
	fmt.Println("p2:", p2)

	p3 := NewPoint(20, 40)
	fmt.Println("p3:", p3)

	r := rect{
		pos:    point{20, 40},
		width:  200,
		height: 400,
	}
	fmt.Println("r:", r)
}
