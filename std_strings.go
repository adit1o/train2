package main

import (
	"fmt"
	"strings"
)

func main() {
    
    adit := "aditya PrataMa"

    fmt.Println(strings.ToTitle(adit))
    fmt.Println(strings.TrimSpace(adit))
    fmt.Println(strings.ReplaceAll(adit, "adit","radit"))
    fmt.Println(strings.ReplaceAll(adit, "tama","dika"))

}