package methods_interfaces

import (
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}


func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.width
}