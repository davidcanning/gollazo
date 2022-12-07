package gollazo

import (
	"fmt"
	"io"
	"os"
)

func SayHello(name string) {
	s := fmt.Sprintf("Hello, %s \n", name)

	io.WriteString(os.Stdout, s)
}
