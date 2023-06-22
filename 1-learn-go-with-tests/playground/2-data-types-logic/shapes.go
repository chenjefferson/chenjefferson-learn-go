package main

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
} // interface resolution is duck typed (implicit)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return
}

func (r Rectangle) Area() (perimeter float64) {
	perimeter = r.Width * r.Height
	return
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() (perimeter float64) {
	perimeter = 2 * math.Pi * c.Radius
	return
}

func (c Circle) Area() (perimeter float64) {
	perimeter = math.Pi * c.Radius * c.Radius
	return
}
