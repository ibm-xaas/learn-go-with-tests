package main

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter ...
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area ...
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

/*
type Circle struct {
	Radius float64
}
*/
func main() {

}
