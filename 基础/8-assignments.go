package main

import "fmt"

func main() {
    var x int
    x = 90
    p := new(bool)
    *p = true
    fmt.Println(x, *p)
}
