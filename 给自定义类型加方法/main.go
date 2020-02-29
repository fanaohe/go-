package main

import "fmt"

// 给自定义类型加方法
type myInt int

func (i myInt) hellp()  {
	fmt.Println("是一个int", i)
	
}

func main()  {
	//var number myInt
	//number = 10
	//number.hellp()

	m := myInt(10)
	m.hellp()
}



