package main

import "fmt"

func main() {
	var p[2]*int
	i := 10
	j := 20
	p[0] = &i
	p[1] = &j

	fmt.Println(p)

	for index ,value := range p{
		fmt.Println(index)
		fmt.Println(value)
		fmt.Println(*value)
	}

}
