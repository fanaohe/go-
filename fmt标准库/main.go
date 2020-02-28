package main

import "fmt"

func main()  {
	a := 10
	fmt.Print("222")    // 默认输出，不加换行
	fmt.Println("222")    // 默认加换行的输出
	fmt.Printf("%T", a)   // print可以格式化输出，可以赋值
	// %T 查看类型
	// %d 十进制
	// %b 二进制
	// %o 八进制
	// %x 十六进制
	// %c 字符
	// %s 字符串
	// %p 指针
	// %v 值   可以在查看变量的值   %#v 查看变量的详细值
	// %f 浮点数
	// %t 布尔类型
	// %c unicode码
	fmt.Printf("%p", &a)    // 返回变量a的16进制指针

	}
