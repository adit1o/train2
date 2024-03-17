package main

/// skippp

import (
	"fmt"
)

type Animal struct {
	Name, Sound string `required:"true" min:"3" max:"10"`
	MaxAge      int    `max:"2"`
}

func main() {
	cat := Animal{
		Name:   "cat",
		Sound:  "meow",
		MaxAge: 10,
	}

	dog := Animal{
		Name:   "dog",
		Sound:  "woof",
		MaxAge: 232,
	}

	fmt.Println(cat)
	fmt.Println(dog)
}
