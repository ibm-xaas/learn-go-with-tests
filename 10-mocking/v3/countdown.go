package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown ...
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprintf(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
