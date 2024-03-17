package main

import (
	"container/list"
	"fmt"
)

func main() {
    data := list.New()

    data.PushBack("adit")
    data.PushBack("tya")
    data.PushBack("pra")
    data.PushBack("tama")


    data.PushFront(" : ")
    data.PushFront("namaku")

    for e := data.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value)
    }

}