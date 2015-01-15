package main

import "fmt"

type point struct {
	x, y int
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
	fmt.Println(p)

	p = point{x: 20, y: 40}
	fmt.Println(p)

	r := rect{
		pos:    point{20, 40},
		width:  200,
		height: 400,
	}
	fmt.Println(r)
}
