package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

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

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
