package printing

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	//now we write our test as greet will accept a buffer and a string and used fprint
	//to allow decoupling
	buffer := bytes.Buffer{}
	Greet(&buffer, "Cheedle")

	got := buffer.String()
	want := "Hello Cheedle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
