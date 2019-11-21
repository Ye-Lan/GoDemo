package main

import "fmt"

func main() {
	nums := [3]int{1,2,3}
	var p *[3]int = &nums
	fmt.Print((*p)[0]) // []的运算优先级高于*
	fmt.Println(p[0])
}
