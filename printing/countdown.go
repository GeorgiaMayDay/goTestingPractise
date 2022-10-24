package printing

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const start = 3

func Countdown(buf io.Writer) {
	for i := start; i > 0; i-- {
		fmt.Fprintf(buf, "%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(buf, finalWord)
}
