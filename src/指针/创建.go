package main

func main() {
	
}

type Student struct {
	id int
	name string
	age int
}

func create1()  {
	s := Student{}
	var student *Student = &s
	print(student)
}

func create()  {
	a := 10
	var num *int = &a
	println(num)
}