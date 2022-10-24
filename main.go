package main

import (
	"fmt"
	"os"
	"time"

	print "goGitTesting/printing"
)

type BaseSleeper struct{}

func Sleep() {
	time.Sleep(1 * time.Second)
}

func Sum(x int, y int) int {
	return x + y
}

func main() {
	sleeper := &print.ConfigurableSleeper{Duration: 1 * time.Second, NewSleep: time.Sleep}
	fmt.Println(Sum(2, 3))
	print.Greet(os.Stdout, "Elodie")
	print.Countdown(os.Stdout, sleeper)
}
