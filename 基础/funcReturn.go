package main

import "fmt"

func main() {
    a := foo()
    fmt.Println(a)
    b := foo2()
    fmt.Println(b)
    aa, bb, cc := foo3()
    fmt.Println(aa, bb, cc)
    a, b, c := foo4()
    fmt.Println(a, b, c)
}

func foo() int {
    return 666
}

// 这种写法很常用
func foo2() (result int) {
    result = 666
    return
}

func foo3() (a int, b int, c string) {
    a, b, c = 111, 222, "abc"
    return a, b, c
}

func foo4() (a, b, c int) {
    a, b, c = 1, 1, 1
    return a, b, c
}
