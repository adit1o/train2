package main

import (
	"flag"
	"fmt"
)
func main() {
    
    host := flag.String("host", "localhost", "put your database host")
    username := flag.String("username", "admin", "put your database username")
    password := flag.String("password", "root", "put your database password")
    port := flag.Int("port", 8080, "put your database port")

    flag.Parse()

    fmt.Println(*host, *username, *password, *port)
}