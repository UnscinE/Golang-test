package main

import (
	"fmt"
	"math"
)

// ประกาศ(Define) Interface ชื่อ Shape
type Shape interface {
	//มีmethod ชื่อ Area() float64
	Area() float64
}

// สร้ําง Struct 2 ตัว
// ชื่อ Rectangle (มี field Width, Height)
type Rectangle struct {
	Width  float64
	Height float64
}

// และ Circle (มี field Radius)
type Circle struct {
	Radius float64
}

func main() {
	//สร้าง instance ของ Rectangle และ Circle
	rect := Rectangle{Width: 5, Height: 4}
	circle := Circle{Radius: 3}

	//แล้วเรียกใช้ฟังก์ชัน PrintArea() เพื่อพิมพ์ขนาดพื้นที่ของทั้ง Rectangle และ Circle ออกมา
	PrintArea(rect)
	PrintArea(circle)

}

// ให้ทั้ง Rectangle และ Circle implement method Area() ตามที่ได้ประกาศไว้ใน Interface Shape
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// เขียนฟังก์ชันแยกออกมา 1 ตัว ชื่อ PrintArea(s Shape) ที่รับ Shape รูปทรงไหนก็ได้
func PrintArea(s Shape) {
	//แล้วพิมพ์ขนาดพื้นที่ออกมา
	fmt.Printf("Area of %T is %.2f\n", s, s.Area())
}
