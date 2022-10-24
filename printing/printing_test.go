package printing

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestGreet(t *testing.T) {
	//now we write our test as greet will accept a buffer and a string and used fprint
	//to allow decoupling
	buffer := bytes.Buffer{}
	Greet(&buffer, "Cheedle")

	got := buffer.String()
	want := "Hello Cheedle \n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleepySpy := SpySleeper{}
	Countdown(&buffer, &sleepySpy)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleepySpy.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", sleepySpy.Calls)
	}
}
