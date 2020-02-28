package main

import "fmt"

// 字符串操作
func main()  {
	//定义字符串
	s1 := "hello你好世界 你好go"
	s2 := []rune(s1)
	changdu := len(s2) - 1
	count := 0
	for i := 0; i<=changdu; i++{
		if s2[i] > 255 {
			count += 1
		}
	}
	fmt.Println(count)
}


