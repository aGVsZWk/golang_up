package main

import "fmt"

func main() {
    foo(10, 20)
    fmt.Printf("\n")
    foo2(10, "abc")
    fmt.Printf("\n")
    foo3(10, 20, 30)
}

func foo(a,b int) {
    fmt.Println(a, b)
}
func foo2(a int, b string) {
    fmt.Println(a, b)
}
func foo3(args ...int){
    for i := range args {
        fmt.Println(args[i])
    }
}
