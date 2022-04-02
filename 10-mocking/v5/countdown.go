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
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// SpySleeper ...
type SpySleeper struct {
	Calls int
}

// Sleep ...
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper ...
type DefaultSleeper struct {
}

// Sleep ...
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
