package main

import "fmt"

func main() {
	num := 2
	test(&num)
	fmt.Println(num)
}

//通过指针避免修改原来的值
func test(num *int){
	num2 := *num + 1
	num = &num2
	fmt.Println(num2)
}
