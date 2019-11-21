package main

import "fmt"

func main() {
	studentMap()
}

type Student struct {
	id int
	name string
	age int
}

func studentMap()  {
	m := make(map[int]Student)
	m[1] = Student{id:1,name:"张三",age:18}
	fmt.Println(m)
}

func students2()  {
	//var arr[]Student = []Student{
	//	Student{id:1,name:"张三",age:18},
	//}
	//arr[0] = Student{id:3,name:"张三",age:18}
	//arr = append(arr, Student{id:2,name:"张三",age:18})
	//fmt.Println(arr)

	arr := make([]Student,5)
	arr[0] = Student{id:3,name:"张三",age:18}
	fmt.Println(arr)
}

func students()  {
	var arr[3] Student = [3]Student{
		Student{id:1,name:"张三",age:18},
		Student{id:2,name:"张三",age:18},
		Student{id:3,name:"张三",age:18},
	}

	fmt.Println(arr)

}

func create()  {
	var s Student = Student{id:1,name:"张三",age:18}
	fmt.Print(s)

	var s1 Student = Student{id:1,name:"张三"}
	fmt.Print(s1)

	var s2 Student
	s2.id = 102
	s2.name = "老王"
	s2.age = 18

	fmt.Println(s2)
}
