package main

import (
    "fmt"
    "time"
)


func main() {
    ch := make(chan int, 0)  //通道无缓冲区
    // ch := make(chan int, 3)  通道缓冲区为 3
    go func() {
        for i := 0; i < 10; i++ {
            // ch <- i
            fmt.Println("goroutine i = ", i)
        }
        time.Sleep(3 * time.Second)
    }()

    for i := 0; i < 10; i++ {
        tmp := <- ch    // 子协程结束，当前只有一个协程（睡了3秒后报的错），直接从管道中读，阻塞个屁！！！直接报错啊！！！
        fmt.Println("main goroutine, chan value = ", tmp)
    }
}
