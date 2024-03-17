package main

import (
	"bufio"
	"os"
)

func main() {

	write := bufio.NewWriter(os.Stdout)
	_, _ = write.WriteString("Hello, ")
	_, _ = write.WriteString("adit here...")

	write.Flush()
}
