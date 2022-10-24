package printing

import (
	"fmt"
	"io"
)

func Greet(buf io.Writer, name string) {
	fmt.Fprintf(buf, "Hello %s", name)
}
