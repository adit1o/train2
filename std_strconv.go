package main

import (
	"fmt"
	"strconv"
)

func main() {
    

    /// string to int
    fmt.Println(strconv.Atoi("12"))

    /// int to string
    fmt.Println(strconv.Itoa(12))


    /// parse
    fmt.Println(strconv.ParseInt("12", 2, 32))

    /// format
    fmt.Println(strconv.FormatInt(999, 2))
    fmt.Println(strconv.FormatBool(false))
}