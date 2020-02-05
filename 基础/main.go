package main

import (
    "calc"      // 执行 calc 的 init()
    "fmt"
)

func init() {    // 先执行自己的 init, 再 main
    fmt.Println("Main init")
}

func main() {
    // t := calc.add(1, 2)
    t := calc.Add(1, 2)
    fmt.Println(t)
}
