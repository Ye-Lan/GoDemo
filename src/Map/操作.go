package main

import "fmt"

func main() {

}

func op3()  {
	m := map[int]string{1:"王五",2:"李四"}

	delete(m , 1)
	printMap(m)
}

func testDeleteMap()  {
	//在函数中修改了map的值会影响原有的Map
	m := map[int]string{1:"王五",2:"李四"}
	deleteMap(m)
	fmt.Println(m)
}

func deleteMap(m map[int]string)  {
	delete(m,2)
	printMap(m)
}

func printMap(m1 map[int]string)  {
	fmt.Print(m1)
}

func op2(){
	m := map[int]string{1:"王五",2:"李四"}
	//key,value := m[1]
	//print(key,value)

	for key,value := range m{
		print(key,value)
	}

}

func op1()  {
	var m map[int] string = map[int]string{1:"王五",2:"李四"}
	value , ok := m[1]
	if ok{
		print(value)
	}else{
		print("不存在")
	}
}
