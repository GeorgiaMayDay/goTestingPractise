package main

import (
	"fmt"
	"os"
	"time"

	print "goGitTesting/printing"
)

type BaseSleeper struct{}

func (s BaseSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Sum(x int, y int) int {
	return x + y
}

func main() {
	sleeper := BaseSleeper{}
	fmt.Println(Sum(2, 3))
	print.Greet(os.Stdout, "Elodie")
	print.Countdown(os.Stdout, sleeper)
}
