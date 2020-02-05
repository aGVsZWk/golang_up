package main

import "fmt"

// 面向过程
func Add01(a, b int) (c int) {
	c = a + b
	return c
}

type long int

// 面向对象：直接给 type 添加方法
func (tmp long) Add02(otherNum long) long {
	return tmp + otherNum
}







func main() {
	var result1, result2 int
	result1 = Add01(10, 20)
	fmt.Println(result1)
	var a long = 10
	// result2 = a.Add02(30) // result2 类型为 int, 不是 long, 报类型不匹配
	r2 := a.Add02(30)
	fmt.Println(r2, result2)
}
