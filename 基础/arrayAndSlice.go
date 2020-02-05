package main

import "fmt"

func main() {
    capAdd()
    a := [3]int {1, 2, 3}
    fmt.Println(a, cap(a), len(a))
    b := []int {1, 2, 3}
    fmt.Println(b, cap(b), len(b))
    c := a[1:2:3]
    fmt.Println(c, cap(c), len(c))
    d := [10]int {}
    fmt.Println(d, len(d), cap(d))
    e := [...]int{1,2,3}
    fmt.Println(e, len(e), cap(e))
    f := make([]int, 3, 5)
    fmt.Println(f, len(f), cap(f))
    f = append(f, 10)
    fmt.Println(f, len(f), cap(f))
    g := make([]int, 10)
    fmt.Println(g, len(g), cap(g))

    a1 := []int {1, 2, 3, 4, 5, 6, 7, 8}
    a2 := a1[1:5]
    a2[3] = 999
    fmt.Println(a1, a2)
    c1:= []int {1, 2, 3}
    c2:= []int {4, 5}
    copy(c1, c2)    // 将 c2 拷贝替换给 c1
    fmt.Println("c1=", c1)
}

// 两倍扩容
func capAdd() {
    s := make([]int, 1)
    oldCap := cap(s)
    for i:=0; i < 8; i++ {
        s = append(s, i)
        if newCap := cap(s); newCap > oldCap {
            fmt.Printf("oldCap=%d ---> newCap=%d\n", oldCap, newCap)
            oldCap = newCap
        }
    }
}
