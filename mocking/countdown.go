package main

import (
	"fmt"
	"io"
)

func main() {

}

func Countdown(writer io.Writer) {
	_, _ = fmt.Fprint(writer, "3")
}
