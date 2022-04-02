package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown ...
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		//time.Sleep(1 * time.Second)
		time.Sleep(1 * time.Microsecond)
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
