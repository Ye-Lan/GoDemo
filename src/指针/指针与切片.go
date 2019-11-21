package main

import "fmt"

func main() {
	s := []int{1,2,3}
	var p *[]int
	p=&s
	fmt.Println(*p)

	//for i:=0; i<len(*p);i++{
	//	fmt.Println((*p)[i])
	//}

	for k,value := range *p{
		fmt.Println("k=",k)
		fmt.Println("value=",value)
	}

}
