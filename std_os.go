package main

import (
	"fmt"
	"os"
)

/// go run std_os.go adit tya pra tama (4)
/// go run std_os.go "adit tya pra tama" (1)

func main() {
    args := os.Args
    for _, arg := range args {
        fmt.Println(arg)
    }

    host, err  := os.Hostname()

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(host)
    }
}