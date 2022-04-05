package main

import "fmt"

func main() {
	s := "Hello 승엽"
	for i, rune := range s {
		fmt.Printf("%d, %#U ", i, rune)
	}
}
