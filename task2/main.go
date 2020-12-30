package main

import "fmt"

type shape interface {
	getArea() float64
}

func main() {

	square := square{sideLength: 5}
	triangle := triangle{base: 5, height: 5}

	printArea(square)
	printArea(triangle)

}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
