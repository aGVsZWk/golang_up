package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 带有接收者的函数叫方法
func (tmp  Person) PrintInfo() {
    fmt.Println("tmp = ", tmp)
}


// 通过一个函数，给成员赋值
func (tmp *Person) SetInfo(n string, s byte, a int) {
    tmp.name = n
    tmp.sex = s
    tmp.age = a
}


// 可给任意自定义类型和内置类型添加方法，除了指针！！！
// type pointer *int
//
// func (tmp  pointer) test() {    // 报错，无效的接受者类型
// }

// 本身不能是指针类型

// func (tmp *int) test() {
//     // ok!!!
// }

// 只要接收类型不一样，函数重名也没关系，不会出现定义重名错误


func main() {
    p := Person{"Mikea", 'm', 22}
    p.PrintInfo()
    var p1 Person
    // p1.SetInfo("John", 'q',100)
    // 上面那行等价于
    (&p1).SetInfo("John", 'q',100)
    p1.PrintInfo()
}
