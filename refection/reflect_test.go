package reflection

import (
	"reflect"
	"testing"
)

type Spy interface {
	Count(string)
	Fetch() []string
}

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name string
	City string
}

type MarriedPerson struct {
	Name     string
	City     string
	Relative Person
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

type NameContainer struct {
	Name      string
	OtherName string
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
		{"struct with string and int field",
			Profile{22, "Daniel"},
			[]string{"Daniel"},
			&SpyWalk{},
		},
		{"struct with with struct with string field",
			MarriedPerson{"Beedle", "Weedle", Person{"Teedle", "Meedle"}},
			[]string{"Beedle", "Weedle", "Teedle", "Meedle"},
			&SpyWalk{},
		},
		{"slice",
			[]Profile{
				{33, "Needle"},
				{44, "Queedle"},
			},
			[]string{"Needle", "Queedle"},
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

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	}
}
