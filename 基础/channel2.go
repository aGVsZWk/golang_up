package main

import (
    "fmt"
    // "time"
)

func main() {
    ch := make(chan int, 0)
    ch <- 10
    go func(tmp chan int){
        fmt.Println(<- tmp)
    }(ch)
}
