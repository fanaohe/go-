package main

import "fmt"

func f1()  {
	fmt.Println("hello word")
}

func f2(x int, y int) int  {
	sum := x + y
	return sum
}

// 函数传参可变参数（y）
func f3(title string, y ...int)int  {
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	return 1
}

// 函数命名返回指
func f4(x, y int) (sum int) {
	sum = x + y   // 如果使用命名返回值 在函数中可以直接进行赋值，无须进行声明
	return    // 在函数中如果有返回值的变量， return后面没有值时可以进行返回，默认时声明的返回变量
}

// 函数进行多命名返回值
func f5(x, y int) (sum, sub int)  {
	sum = x + y
	sub = x - y
	return
}

func main()  {
	f1()
	f1()
	m1 := f2(100, 200)
	fmt.Println(m1)
	f3("title", 1,2,3,4,5)
	fmt.Println(f4(100,500))
	x, y := f5(100, 50)
	m5 := []int{x, y}
	fmt.Println(m5)

}

