package printing

import (
	"bytes"
	"fmt"
)

func Greet(buf *bytes.Buffer, name string) {
	fmt.Println("Hello", name)
}
