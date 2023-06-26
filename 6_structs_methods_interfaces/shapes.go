package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	perimeter := r.Height * r.Width
	return perimeter
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	perimeter := c.Radius * c.Radius * math.Pi
	return perimeter
}

func Perimeter(rectangle Rectangle) float64 {
	perimeter := 2 * (rectangle.Height + rectangle.Width)
	return perimeter
}
