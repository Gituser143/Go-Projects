package main

import "fmt"

type shape interface {
	getArea() float32
	getPerimeter() float32
}

type rect struct {
	width  float32
	length float32
}

type circle struct {
	radius float32
}

func main2() { //Rename to main() when executing
	c := circle{2.0}
	// r := rect{3.0, 5.0}
	printShape(c)
	// printShape(r) // This is an error because it does not have getArea()
}

func printShape(s shape) {
	fmt.Println("I'm a shape")
	fmt.Printf("My perimeter is %v \nMy Area is %v\n", s.getPerimeter(), s.getArea())
}

func (c circle) getArea() float32 {
	return c.radius * c.radius * 3.14
}

func (c circle) getPerimeter() float32 {
	return c.radius * 6.28
}

func (r rect) getPerimeter() float32 {
	return 2 * (r.length + r.width)
}
