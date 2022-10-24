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

type ConfigurableSleeper struct {
	Duration time.Duration
	NewSleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.NewSleep(c.Duration)
}

func Countdown(buf io.Writer, sleep Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintf(buf, "%d\n", i)
		Sleeper.Sleep(sleep)
	}
	fmt.Fprintf(buf, finalWord)
}
