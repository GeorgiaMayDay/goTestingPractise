package printing

import (
	"fmt"
	"io"
)

func Countdown(buf io.Writer) {
	fmt.Fprintf(buf, "3\n")
}
