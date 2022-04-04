package main

import "fmt"

var x int
var y float64

func main() {
	x = 42
	y = 42.3
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
