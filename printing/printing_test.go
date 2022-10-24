package printing

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(dur time.Duration) {
	s.durationSlept = dur
}

type SpyCountdownWatcher struct {
	Operations []string
}

func (s *SpyCountdownWatcher) Write(p []byte) (n int, err error) {
	s.Operations = append(s.Operations, "write")
	return
}

func (s *SpyCountdownWatcher) Sleep() {
	s.Operations = append(s.Operations, "sleep")
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

	t.Run("test content", func(t *testing.T) {
		Countdown(&buffer, &sleepySpy)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if sleepySpy.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", sleepySpy.Calls)
		}
	})

	t.Run("test function calls", func(t *testing.T) {
		fakeWriterAndSleeper := SpyCountdownWatcher{}

		//These need to be passed in as pointers or else
		//the functions that require a pointer can't be called
		Countdown(&fakeWriterAndSleeper, &fakeWriterAndSleeper)

		got := fakeWriterAndSleeper.Operations
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if sleepySpy.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", sleepySpy.Calls)
		}
	})
}

func TestConfigurableSleep(t *testing.T) {
	sleepFor := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{time.Duration(5), spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepFor {
		t.Errorf("should have slept for %v but slept for %v", sleepFor, spyTime.durationSlept)
	}
}
