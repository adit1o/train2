package main

import (
	"errors"
	"fmt"
	"strings"
)



var (
	EmptyError 			= errors.New("empty error")
	ValidationError 	= errors.New("validation error")
	NotFoundError 	    = errors.New("not found error")
)


func GetById (id string) error {

	id = strings.ToLower(id)

	if id == "" {
		return EmptyError
	}
	if id == "anjing" {
		return ValidationError
	}

	if id != "adit" {
		return NotFoundError
	}
	return nil
}

func main() {
	err :=  GetById("asdasd")
	if errors.Is(err, EmptyError) {
		fmt.Println("yee kena `empty` error")
	} else if errors.Is(err, ValidationError) {
		fmt.Println("yee kena `validasi` error")
	} else if errors.Is(err, NotFoundError) {
		fmt.Println("yee kena `notfound` error")
	} else {
		fmt.Println("selesai")
	}
}