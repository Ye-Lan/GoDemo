package main

import "fmt"

func main() {
	testTransfer()
}

/*
切片作为参数，在函数里面修改切片会影响原来切片的值
数组则不会
*/

func testTransfer()  {
	args := []int {0}
	fmt.Println(args)
	transferArgs(args)
	fmt.Println(args)
}

func transferArgs(args []int){
	args[0] = 1
}


// 拷贝的长度为两个切片中长度较小的长度值
func copy2()  {
	s1 := []int{1,2}
	s2 := []int{3,4,5,6}
	copy(s2,s1)
	fmt.Println(s2) // [1 2 5 6]
}


func copy1()  {
	s1 := []int{1,2}
	s2 := []int{3,4,5,6}
	copy(s1,s2)
	fmt.Println(s1) // [3 4]
}


func note(){

	s:=make([]int,5,8)
	fmt.Println(s) //[0 0 0 0 0]
	s = append(s,1)
	fmt.Println(s) //[0 0 0 0 0 1]
	s1 := append(s,2)
	fmt.Println(s1) //[0 0 0 0 0 1 2]
	s1[3] = 33
	fmt.Println(s) //[0 0 0 33 0 1]
	fmt.Println(s1) //[0 0 0 33 0 1 2]
}

func forCycle(){
	//for...len
	s:=[]int{1,2,3,4,5,6,7,8}
	//for i:=0;i<len(s) ; i++  {
	//	fmt.Print(s[i])
	//}

	// for...range
	for _,v := range s{
		fmt.Print(v)
	}
}

func create2(){
	/*
	var a []int
	a = append(a,1,2,3,4,5,6)
	a[3] = 88
 	fmt.Print(a)
	 */

	//s:=[]int{8,9,10,11,12}
	//fmt.Println(s)

	//s := make([]int,3,10)
	//fmt.Print(s) //[0,0,0]
}

func create(){
	// var 切片名 []数据类型
	// 可以使用append函数追加值
	// 注意：切片名[下标]=值，表示的是修改值

	// 切片与数组的区别
	//大小是否可变
	//var b []int = []int{1,2,3}  切片
	//var Numbers[5] int = [5]int{1,2,3,4,5} 数组
	//fmt.Printf("%T,%T",a,b)

	//s:=make([]int,3,5)   //切片类型
	//fmt.Println(s)       //[0 0 0]
	//fmt.Println(len(s))  //3
	//fmt.Println(cap(s))  //5
}
