package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}

func main() {
	stu := Student{0,"admain",10}

	var p *Student = &stu
	fmt.Println(*p)

	test(p)
	fmt.Println(*p)
}

func test(stu *Student)  {
	stu.age = 1
}