package main

import "fmt"

type shape interface {
	area() float64
}

type rect struct {
	length  float64
	breadth float64
}

type square struct {
	side float64
}

func main() {

	r := rect{10, 11}
	s := square{12}

	printArea(r)
	printArea(s)
}

func (r rect) area() float64 {
	return (r.length) * (r.breadth)
}

func (s square) area() float64 {
	return (s.side) * (s.side)
}

func printArea(s shape) {
	fmt.Println(s.area())
}
