package reflection

import (
	"testing"
)

type SpyWalk struct {
	numOfCalls int
}

func (s *SpyWalk) Count(word string) {
	s.numOfCalls += 1
}

func TestWalk(t *testing.T) {
	input := "Cheedle"
	spy := SpyWalk{}

	Walk(input, spy.Count)

	got := spy.numOfCalls
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
