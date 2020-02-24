package main

import "fmt"

// 匿名函数
var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)
	// 普通函数不能用在函数内部，匿名函数可以用在内部
	var f2 = func(x, y int) int {
		//fmt.Println(x + y)
		return x + y
	}

	//在函数内部编写函数，需要直接执行，函数不能有名称，有名称的函数需要编写在外边
	func(x, y int){
		fmt.Println(x + y)
	}(100, 200)

	fmt.Println(f2(30, 50))
}









