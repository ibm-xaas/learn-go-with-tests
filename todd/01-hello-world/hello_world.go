package main

import (
	"fmt"
)

//Hello return a string "Hello, world"
func Hello() string {
	return "Hello, world"
}

func main() {
	Hello()
	fmt.Println(Hello())

}
