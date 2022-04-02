package main

import "math"

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter ...
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func main() {

}
