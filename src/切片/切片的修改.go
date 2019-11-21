package main

import "fmt"

func main() {
	updateValue()
}

func updateValue()  {
	s := []int{1,2,3,4,5,6,7,8}
	s1 := s[2:5]
	//对截取切片值的修改会改变原来切片
	s1[0] = 80
	fmt.Println("s1=",s1) // s1= [80 4 5]
	fmt.Println("s=",s)  // s= [1 2 80 4 5 6 7 8]
}