package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}

func (p Point) Test() float64 {
	return p.x + p.y
}

func pointer_arg(v *Point) float64 {
	return v.x + v.y
}

func (p *Point) TestPointer() float64 {
	p.x += 1
	p.y += 2
	return p.x + p.y
}

func main() {
	p := Point{1, 2}
	fmt.Println(p.Test())
	fmt.Println(p.Test()) // we work with copies so you're not adjusting the same object

	fmt.Println(p.TestPointer())
	fmt.Println(p.TestPointer())
	fmt.Println(pointer_arg(&p)) // When you pas object as ref to function you have to pass it via &
	// What we learned here is that go is explicit with pointers. If no pointer defined we are copying. Look up the exceptions
}
