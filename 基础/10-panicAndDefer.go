package main

import "fmt"

func main() {
    t := recoverFromFor()
    fmt.Println(t)
    // recoverFromBody()
}

func recoverFromFor() int{
    defer func() {
        if msg:= recover(); msg != nil {
            fmt.Println(msg)
        }
    }()

    for i := 1; i <= 10; i++ {
        if i == 5{
            panic("panic...")
        }
        fmt.Println(i)
    }
    fmt.Println("end")
    b := 9
    return b
}

func recoverFromBody() {
    defer func() {
        msg := recover()
        fmt.Println(msg)
    }()
    panic("恐慌...")
    fmt.Println("恢复之后")
}
