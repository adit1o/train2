package main

import (
	"container/ring"
	"fmt"
	"strconv"
)


func main() {
    
    data := ring.New(5)

    for i := 0; i < data.Len(); i++ {
        data.Value = "Data ke " + strconv.Itoa(i)
        data =  data.Next()
    }

    data.Do(func(i any) {
        fmt.Println(i)
    })
}