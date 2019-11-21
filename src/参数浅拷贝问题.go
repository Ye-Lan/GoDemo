package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}

func createStudent() Student  {
	return Student{name:"test",age:1}
}

/*
range本质是一个函数方法，使用时候可以加括号使用
修改range得到后的value,不影响原始切片或者数组
*/

func main() {
	stus := []Student{createStudent(),createStudent(),createStudent()}
	fmt.Println(stus)
	for _,value := range stus{
		value.id = 1
	}
	fmt.Println(stus)
}
