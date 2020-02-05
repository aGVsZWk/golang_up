// type：声明类型
// 使用方法：type 类型名字 底层类型


package main

import "fmt"

type T1 float64
type T2 float64

func main() {
    const (
        t1 T1 = -271.1
        t2 T2 = 0
    )

    fmt.Println(t1, t2)
    fmt.Println(getSumByT1(10, 20))
    fmt.Println(getSumByT2(10, 20))
    fmt.Println(getSumByT2(10, 20) == 30)  // 能和底层自带类型进行比较
    // fmt.Println(getSumByT1(10, 20) == getSumByT2(10, 20)) // 报错！自定义类型不同，不能进行比较
}

func getSumByT1(a float64, b float64) T1 {      // 接收两个 float64，返回自定义的 T1 类型
    return T1(a + b)   // 将 a + b 转换为 T1 类型
}

func getSumByT2(a float64, b float64) T2 {
    return T2(a + b)
}
