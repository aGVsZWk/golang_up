// var ---- variable 变量

package main

import (
    "fmt"
    "os"
)

func main()  {
    var s string
    var i, j, k int     // 默认 0
    var p, q, r = 0, 0, 0.0
    var names []string    // 字符串数组
    var nums []int
    fmt.Println(s, i, j, k, p, q, r, names, nums)


    // 海象运算符 :=
    // 空：nil
    msg := "error"
    f, err := os.Open("a.txt")  // 打开出错, f 为 nil, err 不为 nil, 否则则相反
    if err != nil {
        fmt.Println(msg, err)
    }
    fmt.Println(f)


}
