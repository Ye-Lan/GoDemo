package main

import "fmt"

func main() {

}

func dimensional2Arrag(){
	//var arr[2][3] int
	//arr[0][0] = 1

	//var arr[2] [3] int = [2][3] int {{1,2},{3,4}}
	//fmt.Print(arr)

	//var arr[2] [3] int = [2][3] int {0:{5,6},1:{3,4}}
	//fmt.Print(arr)

	//arr := [...][3]int{{1,2,3},{5,6}}
	//fmt.Print(arr)

	var arr[2] [3] int = [2][3] int {{1,2},{3,4}}
	fmt.Print(len(arr)) // 2 输出结果是几行
	fmt.Print(len(arr[0])) //3 输出结果是几列
	fmt.Print(arr[0][0]) // 1 
}

func argsArray(n [5] int)  {
	for v := range n{
		fmt.Print(v)
	}
}

func cycleArrray() {
	var Numbers [5] int = [5]int{1, 2, 3, 4, 5}

	//for i := 0; i < len(Numbers); i++ {
	//	fmt.Println(Numbers[i])
	//}

	//for i, v := range Numbers {
	//	fmt.Print("下标: ",i)
	//	fmt.Print("z值: ",v)
	//}

	for v := range Numbers {
		fmt.Print("z值: ",v)
	}

}

func initArray() {

	//var Numbers[5] int = [5]int{1,2,3,4,5}
	//fmt.Print(Numbers)
	//[1 2 3 4 5]

	//Numbers := [5]int{1,2}
	//fmt.Print(Numbers)
	//[1 2 0 0 0]

	//Numbers := [5]int{2:5,3:6}
	//fmt.Println(Numbers[2])
	//fmt.Println(Numbers[3])
	// 5 6

	//Numbers := [...]int{7,8,9}
	//fmt.Println(len(Numbers))
	// 3

	//var Numbers[5] int
	//Numbers[0] = 1
	//Numbers[1] = 2
	//fmt.Print(Numbers)
	//[1 2 0 0 0]
}
