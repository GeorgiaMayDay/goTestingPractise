package reflection

import (
	"testing"
)

type Spy interface {
	Count(string)
	Fetch() []string
}

type SpyWalk struct {
	numOfCalls []string
}

func (s *SpyWalk) Count(word string) {
	s.numOfCalls = append(s.numOfCalls, word)
}

func (s *SpyWalk) Fetch() (calls []string) {
	return s.numOfCalls
}

func TestWalk(t *testing.T) {

	cases := []struct {
		name          string
		input         interface{}
		expectedCalls []string
		spy           Spy
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Beedle"},
			[]string{"Beedle"},
			&SpyWalk{},
		},
		{"Just a string",
			"Cheedle",
			[]string{"Cheedle"},
			&SpyWalk{},
		},
		{"struct with two string field",
			struct {
				Name      string
				OtherName string
			}{"Beedle", "Weedle"},
			[]string{"Beedle", "Weedle"},
			&SpyWalk{},
		},
	}

	for _, ex := range cases {

		Walk(ex.input, ex.spy.Count)

		got := ex.spy.Fetch()
		want := ex.expectedCalls

		if len(got) != len(want) {
			t.Errorf("got %d want %d", len(got), len(want))
		}

		if got[0] != want[0] {
			t.Errorf("got %q, want %q", got[0], want[0])
		}
	}
}
