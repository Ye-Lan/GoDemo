package main

import "fmt"

func main() {
	create3()
}

func create3()  {
	m2 := make(map[string] int , 10)
	fmt.Println(m2)
	fmt.Println(len(m2)) // 0
	m2["test"] = 2
	fmt.Println(m2) // map[test:2]

}

func create2(){
	m1 := map[int] string{ 1:"张三",2:"李四",3:"王五"}
	m1[2] = "test"
	fmt.Println(m1)
}

func create(){
	var m map[int] string
	fmt.Print(m)
}
