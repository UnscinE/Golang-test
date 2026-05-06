package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func PrintArea(s Shape) {
	fmt.Printf("Area of %T is %.2f\n", s, s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 4}
	circle := Circle{Radius: 3}

	PrintArea(rect)
	PrintArea(circle)

}
