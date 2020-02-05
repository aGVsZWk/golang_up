package calc

import "fmt"

func init() {       // 包的 init
    fmt.Println("Calc init")
}

// func add(a, b int)int {  # 首字母小写无法被的文件调用，必须首字母大写
func Add(a, b int)int {
    return a + b
}
