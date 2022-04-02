package main

import (
	"bytes"
	"fmt"
)

// Greet ...
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {}
