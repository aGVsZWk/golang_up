package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	id          int
	classId     int
	studyNumber int
	name        string
}

func main() {
	s1 := Student{Person: Person{name: "zhangsan", age: 18}, id: 19, classId: 15, studyNumber: 20133} // 前面的 Person 不能丢
	fmt.Println(s1)
	var s2 Student
	s2.name = "lisi" // 继承！！！就近原则。首先找本作用域，其次作用域链
	s2.age = 18
	fmt.Println(s2)
	s3 := Student{Person{"mike", 18}, 13, 14, 15, "json"}
	fmt.Printf("s3 = %v\n", s3)
	fmt.Println(s3)
	s3.Person = Person{"mike", 22}
	fmt.Println(s3)
	main2()
}

type Student2 struct {
	*Person
	id          int
	classId     int
	studyNumber int
	name        string
}

func main2() {
	s1 := Student2{&Person{"mike", 19}, 13, 14, 15, "json"}
	fmt.Println(s1)
	fmt.Println(s1.age)  // 也可以和之前非指针那样取值

	// 指针定义
	var s2  Student2
	s2.Person = new(Person)
	s2.name = "luke"
	s2.age = 19
	fmt.Println(s2)
}
