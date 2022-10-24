package printing

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const start = 3

type Sleeper interface {
	Sleep()
}

func Countdown(buf io.Writer, sleep Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintf(buf, "%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(buf, finalWord)
}
