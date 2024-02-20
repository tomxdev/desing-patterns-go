package main

import "fmt"

// Shape Prototype Interface
type Shape interface {
	print()
	clone() Shape
}

// Circle Concrete Prototype
type Circle struct {
	color string
}

func (c *Circle) print() {
	fmt.Println(c.color + " circle")
}

func (c *Circle) clone() Shape {
	fmt.Println("Cloning...")
	return &Circle{color: c.color}
}

func main() {
	greenCircle := &Circle{color: "Green"}
	blueCircle := &Circle{color: "Blue"}
	redCircle := &Circle{color: "Red"}

	greenClone := greenCircle.clone()
	blueClone := blueCircle.clone()
	redClone := redCircle.clone()

	greenCircle.print()
	blueCircle.print()
	redCircle.print()

	greenClone.print()
	blueClone.print()
	redClone.print()
}
