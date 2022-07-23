package main

import (
	"bufio"
	"os"

	"golang.org/x/example/stringutil"
)

func main() {
	initialSting := "Hello, OTUS!"
	resultString := stringutil.Reverse(initialSting)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	w.Write([]byte(resultString + "\n"))
}
