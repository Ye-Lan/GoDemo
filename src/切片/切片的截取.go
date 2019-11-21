package main

import (
	"fmt"
)

func main() {
	cutout()
}

func cutout()  {
	s:=[]int{1,2,4,3,5,6,7,8}
	//s[n] 切片中索引位置为n的项
	//第一个值：截取的起始位置
	//第二个值：截取的种植位置（不包含该值）
	//第三个值：用来计算容量，容量指的是切片中最多能够容纳多少元素
	//s1 := s[0:4:7]
	//fmt.Print(s1) //[1,2,3]

	//s1 := s[:]
	//fmt.Println(s1)
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))
	//[1 2 3 4 5 6 7 8]
	//8
	//8

	//s1 := s[3:] //从值为3开始截取
	//fmt.Println(s1) [3 5 6 7 8]
	//fmt.Println(len(s1)) 5
	//fmt.Println(cap(s1)) 5

	//s1 := s[:3]  //截取到3为止
	//fmt.Println(s1) [1 2 4]
	//fmt.Println(len(s1)) 3 len = high
	//fmt.Println(cap(s1)) 8



	s1 := s[1:3]
	fmt.Println(s1) //[2,4]
	fmt.Println(len(s1)) //2
	fmt.Println(cap(s1)) //7 cap = max - low (8-1)

}
