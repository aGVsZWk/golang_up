// 声明

package main

import "fmt"

const boilingF = 202.0   // 常量声明. 在一级范围内声明，可在其它任何地方使用

func main()  {
    var f = boilingF
    fmt.Println(f)

    const freezingF = 201.0   // 函数内部声明
    fmt.Println(freezingF)
}
