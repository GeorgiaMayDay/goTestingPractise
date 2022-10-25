package romanNumerials

import "testing"

func TestNumericalConverter(t *testing.T) {

	t.Run("convert 1 to I", func(t *testing.T) {

		got := Converter(1)
		want := "I"

		if got != want {
			t.Errorf("got %q want %q", want, got)
		}

	})

}
