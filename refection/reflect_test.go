package reflection

import (
	"testing"
)

type SpyWalk struct {
	numOfCalls []string
}

func (s *SpyWalk) Count(word string) {
	s.numOfCalls = append(s.numOfCalls, word)
}

func TestWalk(t *testing.T) {
	input := "Cheedle"
	spy := SpyWalk{}

	Walk(input, spy.Count)

	got := spy.numOfCalls
	want := 1

	if len(got) != want {
		t.Errorf("got %d want %d", len(got), want)
	}

	if got[0] != input {
		t.Errorf("got %q, want %q", got[0], input)
	}
}
