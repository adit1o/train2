package main

import (
	"fmt"
	"time"
)

func main() {

	duration1 := 280000 * time.Millisecond

	fmt.Println(duration1.Minutes())
}
