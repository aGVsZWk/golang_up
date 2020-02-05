package main

import (
    "fmt"
    "time"
)

func main(){
    // var ch chan int     // 这样定义的通道，容量为 0，写入时候直接阻塞了！！！
    // ch := make(chan int, 3) 
    ch := make(chan int, 0)  // 加了容量就 ok 了！！！另一种方法是先消耗至阻塞，后写入，将go协程提前调用！！！

    go func(tmp chan int) {
        fmt.Print(<-tmp)
    }(ch)

    ch <- 10
    // go func(tmp chan int) {
    //     fmt.Print(<-tmp)
    // }(ch)
    time.Sleep(time.Second * 3)
}
