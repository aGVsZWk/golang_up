package main

import "fmt"

type Humaner interface {  // 子集
	sayHi()
}

type Personer interface {  // 超集
	Humaner // 匿名字段，继承了 sayHi
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayHi() {
	fmt.Println("Student say hi", tmp.name, tmp.id)
}

func (tmp * Student) sing(lrc string) {
	fmt.Println("Student sing", lrc)
}

func main() {
	var i Personer
	s := &Student{"Mike", 666}
	i = s
	i.sayHi() // 继承过来的方法
	i.sing("wuawuawua")

}
