package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSorter []User

func (u UserSorter) Len() int {
	return len(u)
}

func (u UserSorter) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{Name: "adit", Age: 20},
		{Name: "budi", Age: 32},
		{Name: "coki", Age: 12},
		{Name: "dede", Age: 18},
	}

	sort.Sort(UserSorter(users))
	fmt.Println(users)
}
