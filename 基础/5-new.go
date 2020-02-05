// new 返回一个地址， new 的本质是一个语法糖，创建一个匿名变量并返回它的地址

package main

import "fmt"

func main() {
    p := new(int)
    fmt.Println(*p)
    p1 := newIntByNew()
    fmt.Println(*p1)
    p2 := newIntByPointer()
    fmt.Println(*p2)
    pointer1 := newIntByNew()
    pointer2 := newIntByNew()
    fmt.Println(pointer1==pointer2)

}


func newIntByNew() *int {
    return new(int)
}


func newIntByPointer() *int {
    var dummy int
    return &dummy
}
