package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown ...
func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
