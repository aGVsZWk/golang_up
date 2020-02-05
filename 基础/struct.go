package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 顺序初始化，每个成员都必须初始化
	var s1 Student = Student{1, "Mike", 'm', 18, "Beijing"}
	fmt.Println("s1=", s1)

	// 指定成员初始化
	s2 := Student{name:"John", addr:"Beijing"}
	s2.age = 30
	fmt.Println(s2)


	// 结构体指针
    var s3 *Student = &Student{3, "John", 'f', 20, "Tangshan"}
    fmt.Println("*s3= ", *s3, "s3=", s3)
	s4 := &Student{name:"John", addr: "Tangshan"}
	fmt.Println(*s4)

	// 指针操作结构体
	p1 := &Student{name: "zhangsan", age: 18}
	// 方式1：(*p1).addr 和 p1.addr 完全等价
	p1.addr = "Beijing"
	(*p1).id = 25
	fmt.Println("*p1=", *p1)

	// 通过 new 创建指针
	p2 := new(Student)
	p2.name = "Luke"
	fmt.Println(*p2)

	// 结构体赋值
	var tmp Student
	tmp = *p2
	fmt.Println(tmp)

	// 结构体值传递 无法改变
	valuePass(tmp)
	fmt.Println("value pass after", tmp)

	// 结构体引用传递，可改变

	referPass(&tmp)
	fmt.Println("refer pass after", tmp)
}

func valuePass(s  Student) {
	s.age = 666
	fmt.Println(s)
	}


func referPass(p *Student) {
	p.age = 666
}
