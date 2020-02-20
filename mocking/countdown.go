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

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(time.Second)
	}
	_, _ = fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
