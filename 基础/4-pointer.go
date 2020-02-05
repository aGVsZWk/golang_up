package main

import "fmt"

func main() {
    x := 1
    p := &x
    fmt.Println(*p)
    *p = 2
    fmt.Println(x)


    // 函数可返回局部变量的地址
    p = f()
    fmt.Println(p)
    fmt.Println(f() == f())     // 每次返回地址都不相同

    *p++  // 只是增加 p 指向的变量的值，并不改变 p 指针！！！

}

func f() *int {
    i := 1
    return &i
}
