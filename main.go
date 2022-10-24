package main

import (
	"fmt"
	"os"

	print "goGitTesting/printing"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Sum(2, 3))
	print.Greet(os.Stdout, "Elodie")
	print.Countdown(os.Stdout)
}
