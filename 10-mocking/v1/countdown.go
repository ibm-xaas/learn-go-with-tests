package main

import (
	"bytes"
	"fmt"
)

// Countdown ...
func Countdown(buffer *bytes.Buffer) {
	fmt.Fprint(buffer, "3")
}

func main() {}
