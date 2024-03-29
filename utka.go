package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type ShapeWithArea interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{3}

	printShapeArea(square)
	printShapeArea(circle)

	printInterface(square)
	printInterface(circle)
	printInterface("qwe")
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())

}

func printInterface(i interface{}) {
	str := i.(string)
	fmt.Println(str)
}
