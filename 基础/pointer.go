package main

import "fmt"

func main() {
    var t int = 10
    var p *int
    fmt.Println(p, "\n")
    // *p = 12 error, p 没有合法指向
    // fmt.Println(*p, "\n")

    p = &t
    fmt.Println(p)
    fmt.Printf("p=%v, &t=%v\n", p, &t)
    *p = 666  // *p 操作的不是p的内存，而是 p 所指向的内存
    fmt.Println(*p, t)
}
