
package main

import "fmt"

type Person struct {
    name  string
    age   int
}

type Student struct {
    Person
    studyId  int
}

// func 定义的方法可继承过去
func (tmp *Person) Study(n string, a int) {
    tmp.name = n
    tmp.age = a
}

// Student 也可重写 Study 方法，重写子类

func (tmp *Student) Study(n string, a, sid int){
    tmp.name = n
    tmp.age = a
    tmp.studyId = sid
}

func main() {
    var p1 Student

    // p1.Study("Mike", 18)     // 调用继承的方法
    // p1.studyId = 3333

    p1.Study("Mike", 18, 333)   // 调用重写的方法
    fmt.Println(p1)

}
