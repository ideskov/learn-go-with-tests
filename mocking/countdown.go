package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout, DefaultSleeper{})
}
