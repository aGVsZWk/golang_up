package main

import "fmt"

func main() {
    // map 类似哈希表，字典
    var m1 map [int] string
    fmt.Println(m1)
    m2 := make(map [int] string) // 无 cap，可指定 len
    // m2 := make(map [int] string, 3) // 无 cap，可指定 len
    fmt.Println(m2)
    m2[1] = "a"
    m2[2] = "b"
    fmt.Println(m2)
    fmt.Println(len(m2))
    m3 := map [int] string{1:"python", 2:"java", 3:"golang"}
    fmt.Println(m3)
    for key, value := range m3{
        fmt.Printf("key=%d, value=%s\n", key, value)
    }
    // 判断一个值是否存在
    // 第一个返回值为 key 所对应的 value，第二个返回值为 key 是否存在的状态 true 或 false
    value, ok := m3[4]
    if ok{
        fmt.Println("key=",value)
    } else {
        fmt.Println("key is not exist")
    }
    // map 删除, delete 指定 key 即可。
    delete(m3, 2)
    fmt.Println(m3)

}
