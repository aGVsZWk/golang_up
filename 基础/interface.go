package main

import "fmt"

type Humaner interface {
    sayHi()
}

type Student struct {
    name  string
    id  int
}

type Teacher struct {
    addr  string
    group string
}

func (s *Student) sayHi() {
    fmt.Println("Hi, I am student", s.name, s.id)
}

func (t *Teacher) sayHi() {
    fmt.Println("Hi, I am teacher", t.addr, t.group)
}

type IntPointer int

func (i *IntPointer) sayHi() {     // 不能写这种指针
    fmt.Println("Hi, I am int")
}

// 多态 定义一个函数，函数的参数为接口的类型
func WhoSayHi(i Humaner) {
    i.sayHi()
}


func main()  {
    // 定义接口类型变量
    var i Humaner

    // 只要定义该接口类型，那么这个类型的变量（接受者类型）就可以给 i 赋值
    s := &Student{"Mike", 3}
    i = s
    i.sayHi()

    t := &Teacher{"Beijing", "g1"}
    i = t
    i.sayHi()

    var n IntPointer = 100
    i = &n
    i.sayHi()


    WhoSayHi(s)
    WhoSayHi(t)
    WhoSayHi(&n)
    x := make([]Humaner, 3)
    x[0] = new(Student)
    x[1] = t
    x[2] = new(IntPointer)
    for _, i = range x {
        WhoSayHi(i)
    }
}
